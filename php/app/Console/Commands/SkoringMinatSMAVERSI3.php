<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use DB;

class SkoringMinatSMAVERSI3 extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'skoring:peminatan_sma_v3';

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

        SKALA_PEMINATAN_SMA                 >>> skoring_minat_sma_v3
        SIKAP_TERHADAP_PELAJARAN            >>> SKORING_SIKAP_PELAJARAN
        SKALA_TES_KARAKTERISTIK_PRIBADI     >>> SKORING_KARAKTERISITIK_PRIBADI

     * @return mixed
     */
    private $space1 = "---> ";
    private $space2 = "--------> ";
    private $space3 = "-------------> ";
    private $paket_soal = 'NON';
    private $tabel_skoring_induk = 'skoring_minat_sma_v3';
    private $tabel_referensi_pilihan_minat = 'ref_pilihan_minat_sma';
    private $tabel_referensi_rekomendasi_pilihan_minat = 'ref_rekomendasi_minat_sma';
    private $tabel_referensi_rekomendasi_akhir = 'ref_rekomendasi_akhir_sma_v3';
    private $kategori_skala_minat = 'SKALA_PEMINATAN_SMA';
    private $id_quiz_template = '1'; //TEMPLATE QUIZ PEMINATAN SMA =>  1; tabel: quiz_sesi_template

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
                    DB::table("$tabel_skoring_induk")->where('id_user', $b->id_user)->delete();
                    DB::table("quiz_sesi_user_jawaban")
                                        ->where('id_user', $b->id_user)
                                        ->where('id_quiz', $b->id_quiz)
                                        ->delete();
                }

                $this->generate_tabel_skor_induk($qz->id_quiz);
                $this->konversi_jawaban($qz->id_quiz);
                $this->skoring_kognitif($qz->id_quiz);
                $this->skoring_minat_sma($qz->id_quiz);
                $this->skoring_sikap_pelajaran($qz->id_quiz);
                $this->skoring_kuliah_ipa($qz->id_quiz);
                $this->skoring_kuliah_ips($qz->id_quiz);
                $this->skoring_sekolah_kedinasan($qz->id_quiz);
                $this->skoring_karakteristik_pribadi($qz->id_quiz);
                $this->skoring_gaya_belajar($qz->id_quiz);
                $this->finishing_skoring($qz->id_quiz);
            }
        }

        $finish_at = date("Y-m-d H:i:s");
        if( $jumlah_running_skoring > 0){
            DB::table('running_cronjob')->insert(['nama'=>"skoring:peminatan_sma_v2",'start_at'=>$start_at, 'finish_at'=>$finish_at, 'peserta'=>$jumlah_running_skoring]);
            //truncate table sementara;
            //DB::table('quiz_sesi_user_jawaban')->truncate();
        }
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
        //$id_quiz = $qz->id_quiz;
         //kosongkan dulu jawaban di tabel quiz_sesi_user_jawaban
        // DB::table('quiz_sesi_user_jawaban')
        //         ->where('id_quiz', $id_quiz)
        //         ->delete();

        //STEP 1 . Pindahkan jawaban user ke tabel quiz_sesi_user_jawaban
        $sesi_user = DB::table('quiz_sesi_user')    //ambil jawaban dari android
                    ->where('id_quiz', $id_quiz)
                    ->where('submit', 1)        //sudah submit
                    ->where('skoring',0)
                    ->where('status_hasil', 0)  //hasil belum ada
                    ->get();
        //var_dump($sesi_user);

        // TABEL: quiz_sesi_user_jawaban
        // id_jawaban_peserta  int
        // id_quiz int
        // id_user int
        // kategori    varchar
        // urutan  int
        // jawaban char
        // skor    int
        echo  $this->space1."Mulai Konversi Jawaban User \n";
        foreach($sesi_user as $su){
            $id_user = $su->id_user;
            $jawaban_submit = json_decode(valid_json_string($su->jawaban));

            if($jawaban_submit){
                //var_dump($jawaban_submit);
                //echo $this->space1."Mulai Konversi Jawaban User ".$su->id_user."\n";
                DB::table("quiz_sesi_user_jawaban")
                                        ->where('id_user', $su->id_user)
                                        ->where('id_quiz', $su->id_quiz)
                                        ->delete();
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

    public function skoring_kognitif($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $paket_soal = $this->paket_soal;
        //Step 1. koreksi jawaban kognitif
        $update_skor = DB::select("update quiz_sesi_user_jawaban
                                        set skor = get_skor_jawaban_kognitif(urutan, jawaban, REPLACE(kategori,'KOGNITIF_',''), '$paket_soal' )
                                    where  
                                        SUBSTR(kategori,1,9)='KOGNITIF_' 
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
                                        AND SUBSTR( a.kategori, 1, 9 )= 'KOGNITIF_' 
                                        AND REPLACE ( a.kategori, 'KOGNITIF_', '' ) = b.nama_bidang 
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
                                tpa_iu + tpa_pv + 
                                tpa_pk + tpa_pa +
                                tpa_ps + tpa_pm + 
                                tpa_kt as skor_tpa_iq  
                                from $tabel_skoring_induk 
                                where id_quiz = $id_quiz and selesai_skoring = 0");

        $ref_konversi_iq = DB::table('ref_konversi_iq')->get();
        $arr_iq = [];
        foreach ($ref_konversi_iq as $rq){
            $arr_iq[$rq->skor_x] = $rq->tot_iq;
        }
        foreach ($skor_iq as $si){
            $record = [];
            //echo $si->skor_tpa_iq. "\n";
            $record['tpa_iq'] = $si->skor_tpa_iq;
            $record['skor_iq'] = $arr_iq[$si->skor_tpa_iq];
            DB::table($tabel_skoring_induk)
                ->where('id', $si->id)
                ->update($record);
        }
        
        echo  $this->space1."Berhasil Skoring Kognitif \n";

    }

    //pake rekom
    public function skoring_minat_sma($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori_skala_minat = $this->kategori_skala_minat; //'SKALA_PEMINATAN_SMA';
        $tabel_referensi_pilihan_minat = $this->tabel_referensi_pilihan_minat;
        $tabel_referensi_rekomendasi_pilihan_minat = $this->tabel_referensi_rekomendasi_pilihan_minat;
        echo  $this->space1."Mulai Skoring ".$kategori_skala_minat." \n";

        $skor_kelompok = DB::select("select 
                                        b.id_quiz, 
                                        b.id_user, 
                                        a.field_skoring,
                                        count(b.jawaban) as total
                                        from 
                                        $tabel_referensi_pilihan_minat as a 
                                        left join quiz_sesi_user_jawaban as b 
                                        on a.kd_pilihan = b.jawaban
                                        and b.id_quiz = $id_quiz and b.kategori = '$kategori_skala_minat'
                                    GROUP BY 
                                        b.id_quiz, 
                                        b.id_user, 
                                        a.field_skoring");
        foreach ($skor_kelompok as $sk){
            $record = [];
            $record[$sk->field_skoring] = $sk->total;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sk->id_quiz)
                ->where('id_user', $sk->id_user)
                ->update($record);
        }

        $skor_rekomendasi = DB::select("
                                        SELECT
                                            x.*,
                                            y.rekomendasi
                                            FROM
                                            (
                                                SELECT
                                                id_user,
                                                id_quiz,
                                                minat_sains,
                                                minat_bahasa,
                                                minat_humaniora,
                                                COALESCE ( minat_sains,0)  - 
                                                COALESCE(minat_bahasa,0) - COALESCE(minat_humaniora, 0) 
                                                AS minat_rentang
                                                FROM
                                                $tabel_skoring_induk AS a
                                                WHERE
                                                a.id_quiz = $id_quiz
                                                and a.selesai_skoring = 0
                                            ) AS x
                                        LEFT JOIN $tabel_referensi_rekomendasi_pilihan_minat AS y ON x.minat_rentang = y.perbedaan ");

        foreach ($skor_rekomendasi as $sr){
            $record = [];
            $record['minat_rentang'] = $sr->minat_rentang;
            $record['rekom_minat'] = $sr->rekomendasi;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sr->id_quiz)
                ->where('id_user', $sr->id_user)
                ->update($record);
        }

        echo  $this->space1."Berhasil Skoring ".$kategori_skala_minat." \n";

    }

    //pake rekom
    public function skoring_sikap_pelajaran($id_quiz){

        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $kategori = 'SIKAP_TERHADAP_PELAJARAN';
        echo  $this->space1."Mulai Skoring ".$kategori." \n";

        $update_skor = DB::select("UPDATE  quiz_sesi_user_jawaban 
                                        set skor = SUBSTR(jawaban,1,1)::INTEGER 
                                        + SUBSTR(jawaban,2,1)::INTEGER 
                                        + SUBSTR(jawaban,3,1)::INTEGER
                                    where id_quiz = $id_quiz 
                                        and kategori = '$kategori'");

        $skor_kelompok = DB::select("SELECT 
                                        a.id_quiz, 
                                        a.id_user, 
                                        a.skor as total , 
                                        b.field_skoring
                                    FROM quiz_sesi_user_jawaban as a, soal_sikap_pelajaran as b  
                                    where 
                                        a.urutan = b.urutan
                                        and a.id_quiz = $id_quiz
                                        and a.kategori = '$kategori' ");

        foreach ($skor_kelompok as $sk){
            $record = [];
            $record[$sk->field_skoring] = $sk->total;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sk->id_quiz)
                ->where('id_user', $sk->id_user)
                ->update($record);
        }

        $skor_kelompok_ilmu = DB::select("SELECT 
                                            a.id_quiz, 
                                            a.id_user, 
                                            c.field_skoring,
                                            sum(a.skor) as total
                                            FROM quiz_sesi_user_jawaban as a, 
                                            soal_sikap_pelajaran as b  , 
                                            ref_kelompok_sikap_pelajaran as c 
                                        where 
                                            a.urutan = b.urutan and a.id_quiz = $id_quiz
                                            and c.kelompok::VARCHAR = b.kelompok::VARCHAR 
                                            and a.kategori = '$kategori'
                                        group by 
                                            id_quiz, id_user, c.field_skoring");

        foreach ($skor_kelompok_ilmu as $ski){
            $record = [];
            $record[$ski->field_skoring] = $ski->total;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $ski->id_quiz)
                ->where('id_user', $ski->id_user)
                ->update($record);
        }


        $skor_rekomendasi = DB::select(" select 
                                            a.id_user, a.id_quiz,
                                            a.sikap_ilmu_alam, 
                                            a.sikap_ilmu_sosial, 
                                            a.sikap_ilmu_alam -  a.sikap_ilmu_sosial as sikap_rentang,
                                            b.rekomendasi 
                                        from $tabel_skoring_induk as a, ref_rekomendasi_sikap_pelajaran as b  
                                        where 
                                            a.id_quiz = $id_quiz 
                                            and b.perbedaan = (a.sikap_ilmu_alam -  a.sikap_ilmu_sosial) 
                                            and a.selesai_skoring = 0
                                            ");

        foreach ($skor_rekomendasi as $sr){
            $record = [];
            $record['sikap_rentang'] = $sr->sikap_rentang;
            $record['rekom_sikap_pelajaran'] = $sr->rekomendasi;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sr->id_quiz)
                ->where('id_user', $sr->id_user)
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

    //rekomendasi akhir
    public function finishing_skoring($id_quiz) {
        $kategori = "AKHIR";
        $tabel_skoring_induk = $this->tabel_skoring_induk;
        $tabel_referensi_rekomendasi_akhir = $this->tabel_referensi_rekomendasi_akhir;

        echo  $this->space1."Mulai Skoring ".$kategori." \n";

        $skor_rekomendasi = DB::select(" 
                                            select a.id_quiz, 
                                                a.id_user, 
                                                a.rekom_minat, 
                                                a.rekom_sikap_pelajaran,
                                                b.rekom_akhir 
                                            from $tabel_skoring_induk as a 
                                            left join $tabel_referensi_rekomendasi_akhir as b
                                                 on a.rekom_minat = b.rekom_minat  
                                                 and a.rekom_sikap_pelajaran = b.rekom_sikap_pelajaran 
                                            where 
                                                a.id_quiz = $id_quiz and a.selesai_skoring = 0
                                             ");

        foreach ($skor_rekomendasi as $sr){
            $record = [];
            $record['rekom_akhir'] = $sr->rekom_akhir;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $sr->id_quiz)
                ->where('id_user', $sr->id_user)
                ->update($record);

        }

        //SKORING KELAS REKOMENDASI
        $skor_kelas = DB::select("SELECT A.id_quiz,
                                A.id_user,
                                A.rekom_minat,
                                A.rekom_sikap_pelajaran,
                                A.rekom_akhir,
                                F.akronim,
                                A.minat_ipa1,
                                A.minat_ips1,
                                A.minat_dinas1,
                            CASE
                                    WHEN i.kelas != '-' THEN
                                    i.kelas 
                                    WHEN A.rekom_akhir = 'ILMU_ALAM' THEN
                                    G.kelas 
                                    WHEN A.rekom_akhir = 'ILMU_SOSIAL' THEN
                                    h.kelas 
                                    WHEN A.rekom_akhir = 'ILMU_ALAM/SOSIAL' THEN
                                    G.kelas ELSE'' 
                                END AS rekom_kelas 
                            FROM
                                skoring_minat_sma_v3 AS A, soal_minat_kuliah_eksakta AS d,
                                soal_minat_kuliah_sosial AS e,
                                ref_sekolah_dinas AS f,
                                ref_kelas_minat_sma AS G,
                                ref_kelas_minat_sma AS h,
                                ref_kelas_minat_sma AS i
                            WHERE
                                A.minat_ipa1 = d.urutan 
                                AND A.minat_ips1 = e.urutan 
                                AND A.minat_dinas1 = f.no
                                AND d.id_kelas = G.ID 
                                AND e.id_kelas = h.ID 
                                AND f.id_kelas = i.ID 
                                and A.selesai_skoring = 0
                                AND A.id_quiz = $id_quiz");

        foreach($skor_kelas as $rk){
            $record = [];
            $record['rekom_kelas'] = $rk->rekom_kelas;
            $record['selesai_skoring'] = 1;
            DB::table($tabel_skoring_induk)
                ->where('id_quiz', $rk->id_quiz)
                ->where('id_user', $rk->id_user)
                ->update($record);

            $record2 = [];   
            $saran = DB::table('quiz_template_saran')
                        ->where('skoring_tabel', $tabel_skoring_induk) 
                        ->first();
            if($saran){
                $saran = $saran->isi;
            }else{
                $saran = "";
            }

            $record2['status_hasil'] = 0;
            $record2['skoring_at'] = date('Y-m-d H:i:s');
            $record2['skoring'] = 1;
            $record2['saran'] = $saran;
            $record2['jawaban_skoring'] = $this->konversi_jawaban_skoring($rk->id_quiz, $rk->id_user);

            DB::table('quiz_sesi_user')
                ->where('id_quiz', $rk->id_quiz)
                ->where('id_user', $rk->id_user)
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

        $result = json_encode($result);
        // DB::table('quiz_sesi_user_jawaban')
        //         ->where('id_quiz', $id_quiz)
        //         ->where('id_user', $id_user)
        //         ->delete();
        return $result;
    }




    //public function konversi_jawaban()


}
