<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use DB;

class SkoringSeleksiKaryawan extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'skoring:seleksi_karyawan';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     
        KOGNITIF_INFORMASI_UMUM
        KOGNITIF_PENALARAN_VERBAL
        KOGNITIF_PENALARAN_KUANTITATIF      >>> SKORING_KOGNITF
        KOGNITIF_PENALARAN_ABSTRAK
        KOGNITIF_PENALARAN_SPASIAL
        KOGNITIF_PENGERTIAN_MEKANIK
        KOGNITIF_CEPAT_TELITI

        SKALA_PEMINATAN_SMA                 >>> skoring_minat_sma_v2
        SIKAP_TERHADAP_PELAJARAN            >>> SKORING_SIKAP_PELAJARAN
        SKALA_TES_MINAT_INDONESIA           >>> SKORING_MINAT_TMI
        SKALA_TES_TIPOLOGI_JUNG             >>> SKORING_TIPOLOGI_JUNG
        SKALA_TES_KARAKTERISTIK_PRIBADI     >>> SKORING_KARAKTERISITIK_PRIBADI

     * @return mixed
     */
    private $space1 = "---> ";
    private $space2 = "--------> ";
    private $space3 = "-------------> ";
    private $paket_soal = 'PMK';
    private $tabel_skoring_induk = 'skoring_seleksi_karyawan';

    public function handle()
    {
        $start_at = date("Y-m-d H:i:s");
        $tabel_skoring_induk =$this->tabel_skoring_induk;
        //$quiz = DB::table('quiz_sesi')->where('skoring_tabel', )->get();
        $quiz = DB::select("
                select 
                a.id_quiz, 
                a.nama_sesi, 
                a.token,
                count(b.id_quiz_user) as peserta_submit 
                from 
                quiz_sesi as a, 
                quiz_sesi_user as b
                where 
                a.id_quiz = b.id_quiz
                and b.submit = 1
                and b.skoring = 0
                and a.skoring_tabel = '$tabel_skoring_induk'
                group by
                a.id_quiz, 
                a.nama_sesi,
                a.token
            ");

        $jumlah_running_skoring = 0;
        foreach ($quiz as $qz){
            
            $ada_skoring =  DB::table('quiz_sesi_user')   
                                    ->select('id_quiz','id_user')
                                    ->where('id_quiz', $qz->id_quiz)
                                    ->where('submit', 1)        //sudah submit
                                    ->where('skoring',0)
                                    ->where('status_hasil', 0)  //hasil belum ada
                                    ->count();
            $jumlah_running_skoring +=$ada_skoring;
            echo "Skoring Quiz ".$qz->nama_sesi." - ".$qz->token."\t".$ada_skoring." User \n";
            if($ada_skoring > 0){
                $belum_skoring = DB::table('quiz_sesi_user')
                                ->select('id_user','id_quiz')
                                    ->where('id_quiz', $qz->id_quiz)
                                    ->where('submit', 1)        //sudah submit
                                    ->where('skoring',0)
                                    ->where('status_hasil', 0) 
                                    ->get();
                foreach($belum_skoring as $b){
                    DB::table("$tabel_skoring_induk")->where('id_user', $b->id_user)->delete();
                    DB::table("quiz_sesi_user_jawaban")
                                        ->where('id_user', $b->id_user)
                                        ->where('id_quiz', $b->id_quiz)
                                        ->delete();
                }

                $this->generate_tabel_skor_induk($qz->id_quiz);
                $this->konversi_jawaban($qz->id_quiz);

                $this->skoring_kognitif($qz->id_quiz); 
                $this->skoring_suasana_kerja($qz->id_quiz);
                $this->skala_kecerdasan_majemuk($qz->id_quiz);
                $this->skoring_minat_tmi($qz->id_quiz);
                $this->skoring_tipologi_jung($qz->id_quiz);
                $this->skoring_karakteristik_pribadi($qz->id_quiz);
                $this->finishing_skoring($qz->id_quiz);
            }
        }

        $finish_at = date("Y-m-d H:i:s");
        if( $jumlah_running_skoring > 0){
            DB::table('running_cronjob')->insert(['nama'=>"skoring:penjurusan_kuliah_v2",'start_at'=>$start_at, 'finish_at'=>$finish_at, 'peserta'=>$jumlah_running_skoring]);
            //truncate table sementara;
            //DB::table('quiz_sesi_user_jawaban')->truncate();
        }
    }

    public function skoring_kognitif($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $paket_soal = $this->paket_soal;
        //Step 1. koreksi jawaban kognitif
        $update_skor = DB::select("update quiz_sesi_user_jawaban
                                        set skor = get_skor_jawaban_kognitif(urutan, jawaban, REPLACE(kategori,'KOGNITIF_PMK_',''), '$paket_soal' )
                                    where  
                                        SUBSTR(kategori,1,13)='KOGNITIF_PMK_' 
                                        and id_quiz = $id_quiz");

        //STEP 2 Grouping dan mapping ke tabel skor_minat_sma
        $skor_kelompok = DB::select("SELECT
                                        a.id_quiz,
                                        a.id_user,
                                        b.nama_bidang,
                                        b.field_skoring,
                                        sum( a.skor ) AS total 
                                    FROM
                                        quiz_sesi_user_jawaban AS a,
                                        ref_bidang_kognitif AS b 
                                    WHERE
                                        a.id_quiz = $id_quiz 
                                        AND SUBSTR( a.kategori, 1, 13 )= 'KOGNITIF_PMK_' 
                                        AND REPLACE ( a.kategori, 'KOGNITIF_PMK_', '' ) = b.nama_bidang 
                                    GROUP BY
                                        a.id_quiz,
                                        a.id_user,
                                        b.nama_bidang,
                                        b.field_skoring");

        foreach ($skor_kelompok as $sk){
            $record = [];
            $record[$sk->field_skoring] = $sk->total;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sk->id_quiz)
                ->where('id_user', $sk->id_user)
                ->update($record);
        }

        $skor_iq = DB::select("select 
                                id, 
                                tpa_iu + tpa_pv + tpa_pk + tpa_pa + tpa_pm + tpa_kt as skor_tpa_iq  
                                from $tabel_skoring_induk 
                                where id_quiz = $id_quiz ");

        $ref_konversi_iq = DB::table('ref_konversi_iq_105')->get();
        $arr_iq = [];
        foreach ($ref_konversi_iq as $rq){
            $arr_iq[$rq->skor_x] = $rq->tot_iq;
        }
        foreach ($skor_iq as $si){
            $record = [];
            $record['tpa_iq'] = $si->skor_tpa_iq;
            //echo("KOGINITIF ".$si->skor_tpa_iq);
            $record['skor_iq'] = $arr_iq[$si->skor_tpa_iq];
            DB::table($tabel_skoring_induk)
                ->where('id', $si->id)
                ->update($record);
        }
        
        echo  $this->space1."Berhasil Skoring Kognitif \n";
    }

    //SKALA KECERDASAN MAJEMUK 
    public function skala_kecerdasan_majemuk($id_quiz){
        $tabel_skoring_induk = $this->tabel_skoring_induk;
        
        $kategori = 'SKALA_KECERDASAN_MAJEMUK';
        $skor_maksimal = 5;
        echo  $this->space1."Mulai Skoring ".$kategori." \n";

        $ref_kecerdasan_majemuk = DB::table('ref_kecerdasan_majemuk')
                        ->select('no')->orderby('no')->get();

        $sesi_user = DB::table('quiz_sesi_user')   
                    ->select('id_quiz','id_user')
                    ->where('id_quiz', $id_quiz)
                    ->where('submit', 1)        //sudah submit
                    ->where('skoring',0)
                    ->where('status_hasil', 0)  //hasil belum ada
                    ->get();
        //reset ref_skoring_kecerdasan_majemuk
        foreach($sesi_user as $su){
            $id_user = $su->id_user; 
            $id_quiz = $su->id_quiz;
            DB::table('ref_skoring_kecerdasan_majemuk')
                    ->where('id_quiz', $id_quiz)
                    ->where('id_user', $id_user)
                    ->delete();
            $jawaban_user = [];
            //ambil jawaban user
            $data_jawaban_user = DB::table('quiz_sesi_user_jawaban')
                            ->where('id_quiz', $id_quiz)
                            ->where('id_user', $id_user)
                            ->where('kategori', $kategori)
                            ->orderby('urutan')->get();
            foreach($data_jawaban_user as $ju){
                $jawaban_user[$ju->urutan] = $ju->jawaban;
            }

            //maping jawaban ke sekolah dinas
            foreach($ref_kecerdasan_majemuk as $r){
                //cek index no di jawaban user
                $b1 = strpos($jawaban_user[1], $r->no) + 1;
                $b2 = strpos($jawaban_user[2], $r->no) + 1;
                $b3 = strpos($jawaban_user[3], $r->no) + 1;
                $b4 = strpos($jawaban_user[4], $r->no) + 1;
                $b5 = strpos($jawaban_user[5], $r->no) + 1;
                $b6 = strpos($jawaban_user[6], $r->no) + 1;
                $b7 = strpos($jawaban_user[7], $r->no) + 1;
                $b8 = strpos($jawaban_user[8], $r->no) + 1;
                $b9 = strpos($jawaban_user[9], $r->no) + 1;

                $record = [
                        'id_quiz'=>$id_quiz,
                        'id_user'=>$id_user,
                        'no'=>$r->no,
                        'b1'=>$b1,
                        'b2'=>$b2,
                        'b3'=>$b3,
                        'b4'=>$b4,
                        'b5'=>$b5,
                        'b6'=>$b6,
                        'b7'=>$b7,
                        'b8'=>$b8,
                        'b9'=>$b9,
                        'total'=>$b1 + $b2 + $b3 + $b4 + $b5 + $b6 + $b7 + $b8 + + $b9,
                        'rangking'=>0,
                    ];
                DB::table('ref_skoring_kecerdasan_majemuk')->insert($record);
            }

            //update rangking
            $current_rangking = DB::table('ref_skoring_kecerdasan_majemuk')
                    ->select('id_skoring_kecerdasan_majemuk')
                    ->where('id_quiz', $id_quiz)
                    ->where('id_user', $id_user)
                    ->orderby('total','asc')
                    ->orderby('no','asc')->get();
            $rangking = 1;
            //var_dump($current_rangking);
            foreach ($current_rangking as $cr){
                DB::table('ref_skoring_kecerdasan_majemuk')
                    ->where('id_skoring_kecerdasan_majemuk', $cr->id_skoring_kecerdasan_majemuk)
                    ->update(['rangking'=>$rangking]);
                $rangking++;
            }

            //update ke tabel skoring induk
            $current_rangking = DB::table('ref_skoring_kecerdasan_majemuk')
                    ->select('no')
                    ->where('id_quiz', $id_quiz)
                    ->where('id_user', $id_user)
                    ->orderby('rangking','asc')
                    ->limit($skor_maksimal)
                    ->get();
            //ambil sebanayak $max
            $ke = 1;
            $update_minat = [];
            //var_dump($current_rangking);
            foreach ($current_rangking as $cr){
                $no = $cr->no;
                $update_minat['km_'.$ke] = $no;
                $ke++;
            }
            DB::table($tabel_skoring_induk)
                    ->where('id_quiz', $id_quiz)
                    ->where('id_user', $id_user)
                    ->update($update_minat);
        }
        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
    }

    //SUASANA KERJA
    public function skoring_suasana_kerja ($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_PMK_SUASANA_KERJA';
        echo  $this->space1."Mulai Skoring ".$kategori." \n";
        $sesi_user = DB::table('quiz_sesi_user')   
                        ->select('id_quiz','id_user')
                        ->where('id_quiz', $id_quiz)
                        ->where('submit', 1)        //sudah submit
                        ->where('skoring',0)
                        ->where('status_hasil', 0)  //hasil belum ada
                        ->get();

        foreach($sesi_user as $su){
            $id_user = $su->id_user; 
            $id_quiz = $su->id_quiz;

            $jawaban = DB::table('quiz_sesi_user_jawaban')
                                ->where('id_quiz', $id_quiz)
                                ->where('kategori', $kategori)
                                ->where('id_user', $id_user)
                                ->orderby('urutan','asc')
                                ->get();
            $record = [];
            foreach ($jawaban as $r){
                $jawaban = $r->jawaban; //ubah ke INTEGER SESUAI FORMAT
                $record['suasana_kerja'.$r->urutan] = $jawaban;
            }

            DB::table($tabel_skoring_induk)
                        ->where('id_quiz', $id_quiz)
                        ->where('id_user', $id_user)
                        ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
    }

    //TMI
    public function skoring_minat_tmi($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_TES_MINAT_INDONESIA';
        $skor_maksimal = 7;


        echo  $this->space1."Mulai Skoring ".$kategori." \n";

        $update_skor = DB::select("update quiz_sesi_user_jawaban 
                                        set skor = ($skor_maksimal + 1) - urutan
                                    where 
                                        id_quiz = $id_quiz and kategori = '$kategori';");

        $skor_kelompok = DB::select("   select
                                            a.id_quiz, 
                                            a.id_user,  
                                            c.field_skoring,
                                            sum(a.skor) as total 
                                        from quiz_sesi_user_jawaban as a, soal_tmi as b, ref_kelompok_tmi as c 
                                        where  
                                            a.id_quiz = $id_quiz 
                                            and a.kategori = '$kategori'
                                            and a.jawaban = lpad(b.urutan::TEXT ,2,'0'::TEXT)
                                            and c.kelompok::VARCHAR = b.kelompok::VARCHAR 
                                        GROUP BY 
                                            a.id_quiz, 
                                            a.id_user,  
                                            c.field_skoring");
        foreach ($skor_kelompok as $sk){
            $record = [];
            $record[$sk->field_skoring] = $sk->total;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sk->id_quiz)
                ->where('id_user', $sk->id_user)
                ->update($record);
        }

        $skor_rekomendasi = DB::select(" select 
                                                a.id_user, a.id_quiz,
                                                a.tmi_ilmu_alam -  a.tmi_ilmu_sosial as tmi_rentang,
                                                b.rekomendasi 
                                            from $tabel_skoring_induk as a,
                                                ref_rekomendasi_minat_tmi as b  
                                            where a.id_quiz = $id_quiz
                                                and b.perbedaan = (a.tmi_ilmu_alam -  a.tmi_ilmu_sosial)
                                                and a.selesai_skoring = 0 
                                             ");

        foreach ($skor_rekomendasi as $sr){
            $record = [];
            $record['tmi_rentang'] = $sr->tmi_rentang;
            $record['rekom_tmi'] = $sr->rekomendasi;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sr->id_quiz)
                ->where('id_user', $sr->id_user)
                ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori." \n";

    }
    //MBTI TIPOLOGI JUNG
    public function skoring_tipologi_jung($id_quiz){
        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_TES_TIPOLOGI_JUNG';

        echo  $this->space1."Mulai Skoring ".$kategori." \n";
        $skor_kelompok = DB::select("select x.id_user, 
                                            x.id_quiz, 
                                            y.kolom, 
                                            z.kode, 
                                            x.skor_a as total_a, 
                                            x.skor_b as total_b, 
                                            z.field_skoring_a, 
                                            z.field_skoring_b
                                            from (SELECT
                                                    a.kolom,
                                                    b.id_quiz, 
                                                    b.id_user,
                                                    sum(case when b.jawaban = 'A' then 1 else 0 end) as skor_a,
                                                    sum(case when b.jawaban = 'B' then 1 else 0 end) as skor_b
                                                FROM
                                                    soal_tipologi_jung AS a,
                                                    quiz_sesi_user_jawaban AS b 
                                                WHERE
                                                    a.urutan = b.urutan 
                                                    AND b.kategori = '$kategori' 
                                                    AND b.id_quiz = $id_quiz  
                                                GROUP BY a.kolom, b.id_quiz, b.id_user) as x, 
                                                ref_skoring_tipologi_jung as y,
                                                ref_klasifikasi_tipologi_jung as z 
                                            where 
                                                x.kolom::VARCHAR = y.kolom::VARCHAR
                                                and x.skor_a = y.skor_a
                                                and y.kode_klasifikasi = z.kode
                                            order by 
                                                x.id_user, y.kolom");

         $tipojung_user = [];
         $tipojung_kode = [];
         $id_user = 0;
         foreach ($skor_kelompok as $sk){
            $record = [];
            $record[$sk->field_skoring_a] = $sk->total_a;
            $record[$sk->field_skoring_b] = $sk->total_b;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sk->id_quiz)
                ->where('id_user', $sk->id_user)
                ->update($record);
            if($sk->id_user!=$id_user){
                $id_user = $sk->id_user;
                array_push($tipojung_user, $id_user);
                $tipojung_kode[$id_user] = $sk->kode;
            }else{
                $tipojung_kode[$id_user] .= $sk->kode;
            }
        }

        for($i = 0;$i<count($tipojung_user);$i++){
            $id_user = $tipojung_user[$i];
            $kode    = $tipojung_kode[$id_user];

            $record = [];
            $record['tipojung_kode'] = $kode;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $id_quiz)
                ->where('id_user', $id_user)
                ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
    }

    public function skoring_karakteristik_pribadi($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_TES_KARAKTERISTIK_PRIBADI';

        echo  $this->space1."Mulai Skoring ".$kategori." \n";

        $update_skor = DB::select("update quiz_sesi_user_jawaban
                                        set skor = get_skor_jawaban_karakteristik_pribadi(jawaban)
                                    where 
                                        id_quiz = $id_quiz and kategori='$kategori' ");
        
        $skor_kelompok = DB::select("SELECT 
                                        a.id_quiz, a.id_user, 
                                        c.field_skoring, sum(a.skor) as total
                                    FROM quiz_sesi_user_jawaban as a  , soal_karakteristik_pribadi as b , ref_komponen_karakteristik_pribadi as c 
                                    where 
                                        a.urutan = b.urutan 
                                        and c.id_komponen = b.id_komponen
                                        and  a.id_quiz = $id_quiz
                                        and a.kategori = '$kategori'
                                    group by 
                                        a.id_user, a.id_quiz, c.field_skoring");

        foreach ($skor_kelompok as $sk){
            $record = [];
            $record[$sk->field_skoring] = $sk->total;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sk->id_quiz)
                ->where('id_user', $sk->id_user)
                ->update($record);
        }

        echo  $this->space1."Mulai Skoring ".$kategori." \n";
    }


    public function generate_tabel_skor_induk($id_quiz){
        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $jumlah_user = 0;
        $sesi_user = DB::table('quiz_sesi_user')   
                    ->select('id_quiz','id_user')
                    ->where('id_quiz', $id_quiz)
                    ->where('submit', 1)        //sudah submit
                    ->where('skoring',0)
                    ->where('status_hasil', 0)  //hasil belum ada
                    ->get();
        foreach($sesi_user as $su){
            $exist_skor_induk = DB::table($tabel_skoring_induk)
                                ->where('id_quiz', $id_quiz)
                                ->where('id_user', $su->id_user)
                                ->count();
            if($exist_skor_induk==0){
                $jumlah_user++;
                $data_awal = array(
                    'id_user'=>$su->id_user, 
                    'id_quiz'=>$su->id_quiz,
                );
                DB::table($tabel_skoring_induk)->insert($data_awal);
            }
        }

        echo  $this->space1."Generate Data Awal Skor ". $tabel_skoring_induk ." - ". $jumlah_user." User \n";
    }
 

    public function konversi_jawaban($id_quiz){
        loadHelper('skoring');
        $sesi_user = DB::table('quiz_sesi_user')    //ambil jawaban dari android
                    ->where('id_quiz', $id_quiz)
                    ->where('submit', 1)        //sudah submit
                    ->where('skoring',0)
                    ->where('status_hasil', 0)  //hasil belum ada
                    ->get();
        echo  $this->space1."Mulai Konversi Jawaban User \n";
        foreach($sesi_user as $su){
            $id_user = $su->id_user;
            $jawaban_submit = json_decode(valid_json_string($su->jawaban));

            if($jawaban_submit){
                //var_dump($jawaban_submit);
                //echo $this->space1."Mulai Konversi Jawaban User ".$su->id_user."\n";
                foreach($jawaban_submit as $js){
                    $kategori = $js->kategori;
                    $jawaban = ($js->jawaban);
                     
                    for($urutan=1;$urutan<count($jawaban);$urutan++){
                        $record_jawaban = array(
                            'id_user'=>$id_user,
                            'id_quiz'=>$id_quiz,
                            'kategori'=>$kategori,
                            'urutan'=>$urutan,
                            'jawaban'=>$jawaban[$urutan],
                            'skor'=>0
                        );
                        DB::table('quiz_sesi_user_jawaban')->insert($record_jawaban);
                    }
                    //echo $this->space2."Berhasil Konversi Jawaban: ".$kategori."\n";
                }
                //echo  $this->space1."Berhasil Konversi Jawaban User ".$su->id_user."\n";
                //echo  $this->space1."=======================================================\n";
            }
        }
        echo  $this->space1."Berhasil Konversi Jawaban User \n";
    }
     
    public function finishing_skoring($id_quiz) {
        $kategori = "AKHIR";
        $tabel_skoring_induk = $this->tabel_skoring_induk;
        
        echo  $this->space1."Mulai Skoring ".$kategori." \n";

        $sesi_user = DB::table($tabel_skoring_induk)   
                        ->select('id_quiz','id_user')
                        ->where('id_quiz', $id_quiz)
                        ->where('selesai_skoring',0)
                        ->get();

        foreach ($sesi_user as $sr){

            //update status skoring dan hasil
            $record = [];
            $record['selesai_skoring'] = 1;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sr->id_quiz)
                ->where('id_user', $sr->id_user)
                ->update($record);

            //ambil saran dari template
            $record2 = [];
            $saran = DB::table('quiz_template_saran')
                        ->where('skoring_tabel', $tabel_skoring_induk) 
                        ->first();
            if($saran){
                $saran = $saran->isi;
            }else{
                $saran = "";
            }
            //$record2['status_hasil'] = 1;
            $record2['skoring_at'] = date('Y-m-d H:i:s');
            $record2['skoring'] = 1;
            $record2['saran'] = $saran;
            $record2['jawaban_skoring'] = $this->konversi_jawaban_skoring($sr->id_quiz, $sr->id_user);

            DB::table('quiz_sesi_user')
                ->where('id_quiz', $sr->id_quiz)
                ->where('id_user', $sr->id_user)
                ->update($record2);

           
        }
        
        echo  $this->space1."Selesai Skoring ".$kategori." \n";
    }

    //konversi jawaban dari tabel quiz_user_sesi_jawaban ke format JSON
    public function konversi_jawaban_skoring($id_quiz, $id_user){

        $data = DB::select("select kategori, urutan, jawaban, skor from quiz_sesi_user_jawaban
                        where id_quiz = $id_quiz and id_user = $id_user
                        order by kategori, urutan asc");
        $result = array();
        $kategori = "";
        $temp = [];
        foreach($data as $r){
            if($kategori != $r->kategori){
                $temp = [];
                $kategori = $r->kategori;
                $obj = json_decode(json_encode(['urutan'=>$r->urutan,'jawaban'=>trim($r->jawaban),'skor'=>$r->skor]));
                array_push($temp,$obj);
            }else{
                $obj = json_decode(json_encode(['urutan'=>$r->urutan, 'jawaban'=>trim($r->jawaban),'skor'=>$r->skor]));
                array_push($temp,$obj);
            }

            $result[$kategori] = $temp;
        }

         //hapus tabel jawaban
            // DB::table('quiz_sesi_user_jawaban')
            //     ->where('id_quiz', $id_quiz)
            //     ->where('id_user', $id_user)
            //     ->delete();
        $result = json_encode($result);
        return $result;
    }


 
}
