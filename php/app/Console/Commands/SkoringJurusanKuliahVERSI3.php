<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use DB;

class SkoringJurusanKuliahVERSI3 extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'skoring:penjurusan_kuliah_v3';

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
    private $tabel_skoring_induk = 'skoring_penjurusan_kuliah_v3';

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
         
        //var_dump($quiz);

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
                    //var_dump($b);
                    DB::table("$tabel_skoring_induk")->where('id_user', $b->id_user)->delete();
                    DB::table("quiz_sesi_user_jawaban")
                                        ->where('id_user', $b->id_user)
                                        ->where('id_quiz', $b->id_quiz)
                                        ->delete();
                }

                $this->generate_tabel_skor_induk($qz->id_quiz);
                $this->konversi_jawaban($qz->id_quiz);
                $this->skoring_kognitif($qz->id_quiz); 
                $this->skoring_sikap_pelajaran($qz->id_quiz);
                $this->skoring_kuliah_ipa($qz->id_quiz);
                $this->skoring_kuliah_ips($qz->id_quiz);
                $this->skoring_kuliah_agama($qz->id_quiz);
                $this->skoring_sekolah_kedinasan($qz->id_quiz);
                $this->skoring_gaya_pekerjaan($qz->id_quiz);
                $this->skoring_gaya_belajar($qz->id_quiz);
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

    public function skoring_sikap_pelajaran($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_PMK_SIKAP_PELAJARAN';
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

            $update_skor = DB::select("UPDATE  quiz_sesi_user_jawaban 
                                        set skor = SUBSTR(jawaban,1,1)::INTEGER 
                                        + SUBSTR(jawaban,2,1)::INTEGER 
                                        + SUBSTR(jawaban,3,1)::INTEGER
                                    where id_quiz = $id_quiz 
                                        and kategori = '$kategori'");
            
            // $update_skor = DB::select("UPDATE  quiz_sesi_user_jawaban 
            //                             set skor = CAST(SUBSTR(jawaban,1,1) as UNSIGNED) 
            //                             + CAST(SUBSTR(jawaban,2,1) as UNSIGNED) 
            //                             + CAST(SUBSTR(jawaban,3,1) as UNSIGNED)
            //                         where id_quiz = $id_quiz and id_user=$id_user and kategori = '$kategori'");

            $skor_kelompok = DB::select("SELECT 
                                        a.id_quiz, 
                                        a.id_user, 
                                        a.skor as total , 
                                        b.field_skoring
                                    FROM quiz_sesi_user_jawaban as a, soal_sikap_pelajaran_kuliah as b  
                                    where 
                                        a.urutan = b.urutan
                                        and a.id_quiz = $id_quiz
                                        and a.id_user = $id_user
                                        and a.kategori = '$kategori' ");
            $record = [];
            foreach ($skor_kelompok as $sk){
                $record[$sk->field_skoring] = $sk->total;
            }
            DB::table($tabel_skoring_induk)
                    ->where('id_quiz', $id_quiz)
                    ->where('id_user', $id_user)
                    ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
    }

    

    //rekom kuliah IPA
    public function skoring_kuliah_ipa ($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_PMK_MINAT_ILMU_ALAM';
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
                $jawaban = (int)$r->jawaban; //ubah ke INTEGER SESUAI FORMAT
                $record['minat_ipa'.$r->urutan] = $jawaban;
            }

            DB::table($tabel_skoring_induk)
                        ->where('id_quiz', $id_quiz)
                        ->where('id_user', $id_user)
                        ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
    }

    //rekom kuliah IPS
    public function skoring_kuliah_ips ($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_PMK_MINAT_ILMU_SOSIAL';
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
                $jawaban = (int)$r->jawaban; //ubah ke INTEGER SESUAI FORMAT
                $record['minat_ips'.$r->urutan] = $jawaban;
            }

            DB::table($tabel_skoring_induk)
                        ->where('id_quiz', $id_quiz)
                        ->where('id_user', $id_user)
                        ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
    }

    //rekom kuliah AGAMA
    public function skoring_kuliah_agama ($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_PMK_ILMU_AGAMA';
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
                $jawaban = (int)$r->jawaban; //ubah ke INTEGER SESUAI FORMAT
                $record['minat_agm'.$r->urutan] = $jawaban;
            }

            DB::table($tabel_skoring_induk)
                        ->where('id_quiz', $id_quiz)
                        ->where('id_user', $id_user)
                        ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
    }


    public function skoring_sekolah_kedinasan($id_quiz){
        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_PMK_SEKOLAH_KEDINASAN';
        $skor_maksimal = 3;
        echo  $this->space1."Mulai Skoring ".$kategori." \n";

        $ref_sekolah_dinas = DB::table('ref_sekolah_dinas')
                        ->select('no')->orderby('no')->get();

        $sesi_user = DB::table('quiz_sesi_user')   
                    ->select('id_quiz','id_user')
                    ->where('id_quiz', $id_quiz)
                    ->where('submit', 1)        //sudah submit
                    ->where('skoring',0)
                    ->where('status_hasil', 0)  //hasil belum ada
                    ->get();
        //reset ref_skoring_kuliah_dinas
        foreach($sesi_user as $su){
            $id_user = $su->id_user; 
            $id_quiz = $su->id_quiz;
            DB::table('ref_skoring_kuliah_dinas')
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
            foreach($ref_sekolah_dinas as $r){
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
                DB::table('ref_skoring_kuliah_dinas')->insert($record);
            }

            //update rangking
            $current_rangking = DB::table('ref_skoring_kuliah_dinas')
                    ->select('id_skoring_sekolah_dinas')
                    ->where('id_quiz', $id_quiz)
                    ->where('id_user', $id_user)
                    ->orderby('total','asc')
                    ->orderby('no','asc')->get();
            $rangking = 1;
            //var_dump($current_rangking);
            foreach ($current_rangking as $cr){
                DB::table('ref_skoring_kuliah_dinas')
                    ->where('id_skoring_sekolah_dinas', $cr->id_skoring_sekolah_dinas)
                    ->update(['rangking'=>$rangking]);
                $rangking++;
            }

            //update ke tabel skoring induk
            $current_rangking = DB::table('ref_skoring_kuliah_dinas')
                    ->select('no')
                    ->where('id_quiz', $id_quiz)
                    ->where('id_user', $id_user)
                    ->orderby('rangking','asc')
                    ->limit($skor_maksimal)
                    ->get();
            //ambil sebanayak $max
            $minat_ke = 1;
            $update_minat = [];
            //var_dump($current_rangking);
            foreach ($current_rangking as $cr){
                $no = $cr->no;
                $update_minat['minat_dinas'.$minat_ke] = $no;
                $minat_ke++;
            }
            DB::table($tabel_skoring_induk)
                    ->where('id_quiz', $id_quiz)
                    ->where('id_user', $id_user)
                    ->update($update_minat);
        }
        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
    }

    public function skoring_gaya_pekerjaan($id_quiz){

        $referensi_komponen = DB::table('ref_komponen_gaya_pekerjaan')->get();
        $kode_komponen = array();
        foreach($referensi_komponen as $r){
            $kode_komponen[$r->no] = $r->kode;
        }

        $user_skoring = DB::table($this->tabel_skoring_induk)
                    ->select('id_user', 'id_quiz')
                    ->where('id_quiz', $id_quiz)->where('selesai_skoring',0)->get();
        foreach($user_skoring as $r){
            $id_user = $r->id_user;
            $id_quiz = $r->id_quiz;

            //hapus skoring terdahulu
            DB::table("ref_skoring_gaya_pekerjaan")
                    ->where('id_user', $id_user)
                    ->where('id_quiz', $id_quiz)
                    ->delete();

            $data_skor = DB::select("select c.jawaban, c.c, c.t, c.u, b.*
                    from quiz_sesi_user_jawaban as a, 
                    soal_gaya_pekerjaan as b , 
                    ref_skor_gaya_pekerjaan as c 
                    where a.urutan = b.nomor 
                    and a.id_quiz = $id_quiz 
                    and a.id_user = $id_user
                    and a.kategori='SKALA_GAYA_PEKERJAAN'
                    and c.jawaban = a.jawaban");

            $array_komponen = array("a","b","c","d","e","f","g","h","i","j","k","l");

            $skor_komponen = array();
            foreach($array_komponen as $n){
                $skor_komponen[$n] = 0;
            }
            foreach($data_skor as $s){
                $skor_U = $s->u;
                $skor_T = $s->t;
                $skor_C = $s->c;
                foreach($array_komponen as $n){
                    $nama_komponen = "komponen_".$n;
                    if($s->$nama_komponen!="" && $s->$nama_komponen!=null){
                        $cek = substr($s->$nama_komponen, 0, 1);
                        if($cek=="U"){
                            $skor_komponen[$n] += $skor_U;
                        }
                        if($cek=="T"){
                            $skor_komponen[$n] += $skor_T;
                        }
                        if($cek=="C"){
                            $skor_komponen[$n] += $skor_C;
                        }
                    }
                }
            }

            arsort($skor_komponen);//urutkan rangking
            $record_skoring_induk = array();
            $batas_rangking = 3;
            $rangking = 1;
            foreach($skor_komponen as $key => $value){
                $nama_komponen = "komponen_".$key;
                //echo $kode_komponen[$key] ."\t: ".$value."\n";
                $field_name = "gp_".$key; 
                $record_skoring_induk[$field_name] = $value;

                if($rangking<=$batas_rangking){
                    $field_rangking = "rangking_gp".$rangking;
                    $record_skoring_induk[$field_rangking] = $kode_komponen[$key];
                }

                $record_ref_skoring = array("id_user"=>$id_user, 
                                            "id_quiz"=>$id_quiz, 
                                            "kode"=>$kode_komponen[$key],
                                            "skor"=>$value, 
                                            "rangking"=>$rangking);

                DB::table("ref_skoring_gaya_pekerjaan")->insert($record_ref_skoring);
                $rangking++;
            }
            //var_dump($record_skoring_induk);
            DB::table($this->tabel_skoring_induk)
                    ->where('id_user', $id_user)
                    ->where('id_quiz', $id_quiz)
                    ->update($record_skoring_induk);

            //update klasifikasi skor gaya pekerjaan
            $update_klasifikasi = DB::select("select a.id,   a.skor, b.akronim as klasifikasi
                        from ref_skoring_gaya_pekerjaan  as a, 
                        ref_klasifikasi_gaya_kerja as b  
                        where a.skor >= b.skor_min and a.skor <= b.skor_max  
                        and a.id_quiz = $id_quiz and a.id_user = $id_user
                        order by a.rangking");

            foreach($update_klasifikasi as $u){
                DB::table('ref_skoring_gaya_pekerjaan')
                    ->where('id', $u->id)
                    ->update(['klasifikasi'=>$u->klasifikasi]);
            }
        }

        echo  $this->space1."Berhasil Skoring Gaya Pekerjaan \n";

    }

    public function skoring_gaya_belajar($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SKALA_GAYA_BELAJAR';
 
        $tabel_skoring_induk = $this->tabel_skoring_induk; 
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
            $jawaban = DB::select("SELECT
                                            id_user,
                                            jawaban,
                                            COUNT ( * ) AS jumlah 
                                        FROM
                                            quiz_sesi_user_jawaban 
                                        WHERE
                                            id_quiz = $id_quiz 
                                            AND kategori = '$kategori' 
                                            AND id_user = $id_user
                                        GROUP BY
                                            id_user,
                                            jawaban 
                                        ORDER BY
                                            id_user,
                                            jawaban");
            $record = [];
            $record['gaya_auditoris'] = 0;
            $record['gaya_visual'] = 0;
            $record['gaya_kinestetik'] = 0;

            foreach ($jawaban as $r){
                $jawaban = $r->jawaban;
                if ($jawaban =="A"){
                   $record['gaya_auditoris'] = $r->jumlah; 
                }

                if ($jawaban =="B"){
                   $record['gaya_visual'] = $r->jumlah; 
                }

                if ($jawaban =="C"){
                   $record['gaya_kinestetik'] = $r->jumlah; 
                }
            }

            DB::table($tabel_skoring_induk)
                        ->where('id_quiz', $id_quiz)
                        ->where('id_user', $id_user)
                        ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori." \n";
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
                        // var_dump($record_jawaban);
                        //hapus dulu
                        DB::table('quiz_sesi_user_jawaban')
                                ->where('id_user',$id_user)
                                ->where('id_quiz',$id_quiz)
                                ->where('kategori',$kategori)
                                ->where('urutan',$urutan)
                                ->delete();
                                
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




    //public function konversi_jawaban()


}
