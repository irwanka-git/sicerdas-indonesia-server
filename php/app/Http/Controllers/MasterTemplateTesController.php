<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;

use Symfony\Component\Process\Process;
use Symfony\Component\Process\Exception\ProcessFailedException;  
use setasign\Fpdi\Fpdi;

class MasterTemplateTesController extends Controller
{
     
    //######################### QUIZ SESI MASTER TEMPLATE #####################################
    // quiz_sesi_template
    // id_quiz_template int
    // nama_sesi        varchar
    // skoring_tabel    varchar
    // uuid             char

    // quiz_sesi_detil_template
    // id_quiz_sesi_template   int
    // id_quiz_template int
    // id_sesi_master  int
    // urutan  int
    // durasi  int
    // kunci_waktu int

    function index(){
    	$pagetitle = "Master Template Tes";
    	$smalltitle = "Data Master Template Tes";
    	return view('master-template.index', compact('pagetitle','smalltitle'));
    }

    function datatable_template(Request $r){
        
        $filter = "";
        if (request()->has('search')) {
            $search = request('search');
            $keyword = $search['value'];
            if(strlen($keyword)>=2){
                $keyword = strtolower($keyword);
                $filter = " and (  lower(a.nama_sesi) like '%$keyword%' or lower(a.kode) like '%$keyword%' ) ";
            }   
        }

         $sql_union = "select a.id_quiz_template,  a.kode,
            a.nama_sesi, a.gambar,
            a.uuid, a.jenis, 
            count(b.id_sesi_master) as jumlah_sesi
            from quiz_sesi_template as a 
            left join quiz_sesi_detil_template as b 
                on a.id_quiz_template = b.id_quiz_template
                where a.id_quiz_template > 0 $filter
            group by a.id_quiz_template, 
            a.nama_sesi, a.skoring_tabel, 
            a.uuid 
        order by a.jenis  desc , a.kode asc";
        //return $sql_union;
         $query = DB::table(DB::raw("($sql_union) as x order by jenis desc, kode asc"))
                    ->select([
                        'id_quiz_template',
                        'kode',
                        'nama_sesi', 
                        'gambar',
                        'jumlah_sesi', 
                        'jenis',
                        'uuid',
                    ]);

         return Datatables::of($query)
            ->addColumn('action', function ($query) {
                $edit = ""; $delete = "";
                $action = '';
                if($this->ucu()){

                $edit = '<button data-bs-tip="tooltip" data-bs-placement="top" data-bs-original-title="Ubah Jenis Tes"  data-bs-toggle="modal" data-uuid="'.$query->uuid.'" data-bs-target="#modal-edit" class="btn btn-light btn-sm" type="button"><ion-icon name="create-outline" btn></ion-icon></button>';
                }
                if($this->ucd()){
                    $delete = '<button data-bs-tip="tooltip" data-bs-placement="top" data-bs-original-title="Hapus Jenis Tes"   data-uuid="'.$query->uuid.'" class="btn btn-light btn-sm btn-konfirm-delete" type="button"><ion-icon name="trash-outline" btn></ion-icon></button>';
                }
                $action =  $action." ".$edit." ".$delete;
                if ($action==""){return '<a href="#" class="act"><i class="la la-lock"></i></a>'; }
                return '<div class="btn-group" role="group">'.$action.'</button>';
            })
            ->editColumn('gambar', function ($q) {
                if($q->gambar){
                    $gambar = basename($q->gambar).PHP_EOL;
                  return '<button data-image="'.$q->gambar.'" class="btn btn-outline-secondary btn-outline btn-sm btn-view-image" type="button">'.$gambar.'</button>';
                }
            })
            ->editColumn('jumlah_sesi', function($q){
                // return "<a href='".url('template-tes/detil/'.$q->uuid)."'>". $q->jumlah_sesi ." Sesi</a>";
                return "<a class='btn btn-sm btn-outline-secondary' href='".url('template-tes/detil/'.$q->uuid)."'> <i class='la la-list'></i> ". $q->jumlah_sesi ." Sesi </a>";
            })
            ->editColumn('item_report', function($q){
                if ($q->jenis=="quiz") 
                return "<a class='btn btn-sm btn-outline-secondary' href='".url('template-tes/report/'.$q->uuid)."'> <i class='la la-cog'></i> Laporan </a>";
            })
            ->editColumn('kode', function($q){
                 return $q->kode."-".$q->id_quiz_template;
            })
            ->addIndexColumn()
            ->rawColumns(['action','jumlah_sesi','gambar','item_report'])
            ->make(true);
    }

    function get_data($uuid){
    	$data = DB::table('quiz_sesi_template')->where('uuid', $uuid)->first();
        if($data){
            $respon = array('status'=>true,'data'=>$data, 
            	'informasi'=>'Master Template : '.$data->nama_sesi);
            return response()->json($respon);
        }
        $respon = array('status'=>false,'message'=>'Data Tidak Ditemukan');
        return response()->json($respon);
    }

    function submit_insert(Request $r){
        if($this->ucc()){
            loadHelper('format');
	    	$uuid = $this->genUUID();
            // $petunjuk_sesi = str_replace("<p>", "", $r->petunjuk_sesi);
            // $petunjuk_sesi = str_replace("</p>", "", $petunjuk_sesi);          
            // return $petunjuk_sesi;
	    	$record = array(                                              
                "kode"=>trim($r->kode),
                "nama_sesi"=>trim($r->nama_sesi),
                "gambar"=>trim($r->gambar),
                "jenis"=>trim($r->jenis),
                // "skoring_tabel"=>trim($r->skoring_tabel),
                "pendahuluan"=>trim($r->pendahuluan),
                "kertas"=>"Folio",
	    		"uuid"=>$uuid,
            );

	    	DB::table('quiz_sesi_template')->insert($record);
	    	$respon = array('status'=>true,'message'=>'Data Berhasil Ditambahkan!');
        	return response()->json($respon);
    	}else{
    		$respon = array('status'=>false,'message'=>'Akses Ditolak!');
        	return response()->json($respon);
    	}
    }

    function submit_update(Request $r){
    	if($this->ucu()){
	    	loadHelper('format');
	    	$uuid = $r->uuid;
            // $pernyataan = str_replace("<p>", "", $r->pernyataan);
            // $pernyataan = str_replace("</p>", "", $pernyataan);
            //return $pernyataan;
	    	$record = array(
                "kode"=>trim($r->kode),
                 "nama_sesi"=>trim($r->nama_sesi),
                  "gambar"=>trim($r->gambar),
                  "jenis"=>trim($r->jenis),
                  "pendahuluan"=>trim($r->pendahuluan),
                // "skoring_tabel"=>trim($r->skoring_tabel)
            );

	    	DB::table('quiz_sesi_template')->where('uuid', $uuid)->update($record);
	    	$respon = array('status'=>true,'message'=>'Data Berhasil Disimpan!');
        	return response()->json($respon);
    	}else{
    		$respon = array('status'=>false,'message'=>'Akses Ditolak!');
        	return response()->json($respon);
    	}
    }

    function update_kertas(Request $r){
        if($this->ucu()){
	    	loadHelper('format');
	    	$id_quiz_template = $r->id_quiz_template;
	    	$record = array(
                  "kertas"=>trim($r->kertas),
            );
	    	DB::table('quiz_sesi_template')->where('id_quiz_template', $id_quiz_template)->update($record);
	    	$respon = array('status'=>true,'message'=>'Ukuran Kertas Berhasil Disimpan!');
        	return response()->json($respon);
    	}else{
    		$respon = array('status'=>false,'message'=>'Akses Ditolak!');
        	return response()->json($respon);
    	}
    }

    
    function submit_update_saran(Request $r){
    	if($this->ucu()){
	    	loadHelper('format');
	    	$id = $r->id_quiz_template;
            $isi_saran = trim ($r->isi_saran);
	    	$record = array(
                  "isi_saran"=>trim($r->isi_saran),
                  "judul_saran"=>trim($r->judul_saran),
            );
	    	DB::table('quiz_sesi_template')->where('id_quiz_template', $id)->update($record);
	    	$respon = array('status'=>true,'message'=>'Saran Berhasil Disimpan!');
        	return response()->json($respon);
    	}else{
    		$respon = array('status'=>false,'message'=>'Akses Ditolak!');
        	return response()->json($respon);
    	}
    }

    function submit_delete(Request $r){
        if($this->ucd()){
            loadHelper('format');
            $uuid = $r->uuid;
            DB::table('quiz_sesi_template')->where('uuid', $uuid)->delete();
            $respon = array('status'=>true,'message'=>'Data Berhasil Dihapus!');          
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }
    //#######################################################################################
    //DETIL

    function index_detil($uuid){
        $pagetitle = "Master Template Tes";
        $smalltitle = "Data Master Template Tes";
        $template = DB::table('quiz_sesi_template')->where('uuid', $uuid)->first();
        return view('master-template.index-detil', compact('pagetitle','smalltitle','template'));
    }

    function datatable_detil($uuid){
        $template = DB::table('quiz_sesi_template')->where('uuid', $uuid)->first();
        $id_quiz_template = $template->id_quiz_template;
        $filter = "";
        if (request()->has('search')) {
            $search = request('search');
            $keyword = $search['value'];
            if(strlen($keyword)>=2){
                $keyword = strtolower($keyword);
                $filter = " and (  lower(a.nama_sesi) like '%$keyword%' ) ";
            }   
        }

         $sql_union = "select 
                            a.id_quiz_sesi_template,
                            a.urutan,
                            a.durasi,
                            a.kunci_waktu,
                            b.nama_sesi_ujian, a.uuid
                        FROM
                            quiz_sesi_detil_template AS a,
                            quiz_sesi_master AS b 
                        WHERE
                            a.id_sesi_master = b.id_sesi_master 
                            AND a.id_quiz_template =  $id_quiz_template $filter ";
        //return $sql_union;
         $query = DB::table(DB::raw("($sql_union) as x order by urutan"))
                    ->select([
                        'id_quiz_sesi_template',
                        'urutan',
                        'nama_sesi_ujian',
                        'kunci_waktu',
                        'durasi',
                        'uuid',
                    ]);

         return Datatables::of($query)
            ->addColumn('action', function ($query) {
                    $edit = ""; $delete = "";
                    $action = '';
                    if($this->ucu()){

                    $edit = '<button data-bs-tip="tooltip" data-bs-placement="top" data-bs-original-title="Ubah Sesi"  data-bs-toggle="modal" data-uuid="'.$query->uuid.'" data-bs-target="#modal-edit" class="btn btn-light btn-sm" type="button"><ion-icon name="create-outline" btn></ion-icon></button>';
                    }
                    if($this->ucd()){
                        $delete = '<button data-bs-tip="tooltip" data-bs-placement="top" data-bs-original-title="Hapus Sesi"   data-uuid="'.$query->uuid.'" class="btn btn-light btn-sm btn-konfirm-delete" type="button"><ion-icon name="trash-outline" btn></ion-icon></button>';
                    }
                    $action =  $action." ".$edit." ".$delete;
                    if ($action==""){return '<a href="#" class="act"><i class="la la-lock"></i></a>'; }
                    return '<div class="btn-group" role="group">'.$action.'</button>';
            }) 
            ->rawColumns(['action'])
            ->make(true);
    }

    function get_data_detil($uuid){
       $data = DB::table('quiz_sesi_detil_template')->where('uuid', $uuid)->first();
        if($data){
            $respon = array('status'=>true,'data'=>$data, 
                'informasi'=>'Sesi Urutan : '.$data->urutan);
            return response()->json($respon);
        }
        $respon = array('status'=>false,'message'=>'Data Tidak Ditemukan');
        return response()->json($respon);     
    }

    function submit_insert_detil(Request $r){
        if($this->ucc()){
            loadHelper('format');
            $uuid = $this->genUUID();
            
            $record = array( 
                "id_quiz_template"=>$r->id_quiz_template,                                        
                "urutan"=>(int)($r->urutan),
                "durasi"=>(int)($r->durasi),
                "kunci_waktu"=>(int)($r->kunci_waktu),
                "id_sesi_master"=>(int)($r->id_sesi_master),
                "uuid"=>$uuid,
            );

            DB::table('quiz_sesi_detil_template')->insert($record);
            $respon = array('status'=>true,'message'=>'Data Berhasil Ditambahkan!');
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }

    function submit_update_detil(Request $r){
        if($this->ucu()){
            loadHelper('format');
            //$uuid = $this->genUUID();
            
            $record = array(                                    
                "urutan"=>(int)($r->urutan),
                "durasi"=>(int)($r->durasi),
                "kunci_waktu"=>(int)($r->kunci_waktu),
                "id_sesi_master"=>(int)($r->id_sesi_master),
            );

            DB::table('quiz_sesi_detil_template')->where('uuid', $r->uuid)->update($record);
            $respon = array('status'=>true,'message'=>'Data Berhasil Disimpan!');
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }

    function submit_delete_detil(Request $r){
        if($this->ucu()){
            //loadHelper('format');
            DB::table('quiz_sesi_detil_template')->where('uuid', $r->uuid)->delete();
            $respon = array('status'=>true,'message'=>'Data Berhasil Dihapus!');
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }

    ///
    function index_report($uuid){
        $pagetitle = "Master Template Tes";
        $smalltitle = "Data Report - Template Tes";
        $template = DB::table('quiz_sesi_template')->where('uuid', $uuid)->first();
        $id_quiz_template = $template->id_quiz_template;
        $sesi = DB::select("select 
                                a.id_quiz_sesi_template,
                                b.kategori, 
                                b.tabel,
                                b.nama_sesi_ujian, 
                                a.urutan,
                                a.uuid
                            FROM
                                quiz_sesi_detil_template AS a,
                                quiz_sesi_master AS b 
                            WHERE
                                a.id_sesi_master = b.id_sesi_master 
                                AND a.id_quiz_template =  $id_quiz_template order by a.urutan asc");
        return view('master-template.index-report', compact('pagetitle','smalltitle','template','sesi'));
    }

    function datatable_report($uuid){
        $template = DB::table('quiz_sesi_template')->where('uuid', $uuid)->first();
        $id_quiz_template = $template->id_quiz_template;
        $filter = "";
        if (request()->has('search')) {
            $search = request('search');
            $keyword = $search['value'];
            if(strlen($keyword)>=2){
                $keyword = strtolower($keyword);
                $filter = " and (  lower(a.nama_report) like '%$keyword%' ) ";
            }   
        }

         $sql_union = "select 
                            b.urutan,
                            a.nama_report,
                            b.uuid
                        FROM
                            quiz_sesi_template_report AS b,
                            quiz_sesi_report AS a 
                        WHERE
                            a.id_report = b.id_report 
                            AND b.id_quiz_template =  $id_quiz_template $filter ";
        //return $sql_union;
         $query = DB::table(DB::raw("($sql_union) as x order by urutan"))
                    ->select([
                        'urutan',
                        'nama_report',
                        'uuid',
                    ]);

         return Datatables::of($query)
            ->addColumn('action', function ($query) {
                    $edit = ""; $delete = "";
                    $action = '';
                     
                    if($this->ucd()){
                        $delete = '<button data-bs-tip="tooltip" data-bs-placement="top" data-bs-original-title="Hapus Report"   data-uuid="'.$query->uuid.'" class="btn btn-light btn-sm btn-konfirm-delete" type="button"><ion-icon name="trash-outline" btn></ion-icon></button>';
                    }
                    $action =  $action." ".$edit." ".$delete;
                    if ($action==""){return '<a href="#" class="act"><i class="la la-lock"></i></a>'; }
                    return '<div class="btn-group" role="group">'.$action.'</button>';
            })
            ->addIndexColumn()
            ->rawColumns(['action'])
            ->make(true);
    }

    function submit_insert_report(Request $r){
        if($this->ucc()){
            loadHelper('format');
            $uuid = $this->genUUID();
            $model = trim($r->model);
            $id_quiz_template = $r->id_quiz_template;
            $id_report = (int)($r->id_report);
            $tabel_referensi = DB::table('quiz_sesi_report')->where('id_report', $id_report)->first()->tabel_referensi;
            if($tabel_referensi!="-"){
                $exist = DB::table('quiz_sesi_template_report')
                    ->where('id_quiz_template', $id_quiz_template)
                    ->where('model', $model)
                    ->where('id_report',$id_report)
                    ->count();
                if($exist){
                    $respon = array('status'=>false,'message'=>'Komponen sudah ada!');
                    return response()->json($respon);
                }
            }
           
            $jumlah = DB::select(
                "select 
                count(*) as jumlah
                FROM
                    quiz_sesi_template_report AS b,
                    quiz_sesi_report AS a 
                WHERE
                    a.id_report = b.id_report 
                    AND b.id_quiz_template =  $id_quiz_template"
            );
            $urutan = $jumlah[0]->jumlah + 1;
            $record = array( 
                "id_quiz_template"=>$r->id_quiz_template,                                        
                "urutan"=>(int)($urutan),
                "id_report"=>$id_report,
                "model"=>$model,
                "uuid"=>$uuid,
            );
            DB::table('quiz_sesi_template_report')->insert($record);
            $respon = array('status'=>true,'message'=>'Komponen Berhasil Ditambahkan!');
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }


    function get_data_report($uuid){
        $data = DB::table('quiz_sesi_template_report')->where('uuid', $uuid)->first();
        if($data){
            $report = DB::table('quiz_sesi_report')->where('id_report', $data->id_report)->first();
            $respon = array('status'=>true,'data'=>$data, 
                'informasi'=>'Komponen : '.$report->nama_report);
            return response()->json($respon);
        }
        $respon = array('status'=>false,'message'=>'Data Tidak Ditemukan');
        return response()->json($respon); 
    }

    function submit_delete_report(Request $r){
        if($this->ucu()){
            //loadHelper('format');
            $id_quiz_template = DB::table('quiz_sesi_template_report')->where('uuid', $r->uuid)->first()->id_quiz_template;
            DB::table('quiz_sesi_template_report')->where('uuid', $r->uuid)->delete();
            $listbaru = DB::table('quiz_sesi_template_report')
                            ->where('id_quiz_template', $id_quiz_template)
                            ->orderby("urutan")->get();
            $urutan = 1;
            foreach ($listbaru as $r){
                $uuid = $r->uuid;
                DB::table('quiz_sesi_template_report')->where('uuid', $uuid)->update(["urutan"=>$urutan]);
                $urutan++;
            }
            $respon = array('status'=>true,'message'=>'Komponen Berhasil Dihapus!');
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }

    function get_list_komponen_report($id_quiz_template, $model){
        
        $list = DB::select("select a.nama_report, b.uuid ,
                 b.urutan, a.tabel_referensi 
                        from quiz_sesi_report as a, 
                            quiz_sesi_template_report as b where a.jenis = 1 
                                and a.id_report = b.id_report and b.model = '$model'
                                    and b.id_quiz_template = $id_quiz_template order by b.urutan asc");
        return view('master-template.drag-report', compact('list'));
    }

    function submit_update_urutan(Request $r){
        if($this->ucc()){
            $urutan_string =  $r->urutan_list;
            $urutan_list = explode(",", $urutan_string);
            $urutan = 1;
            foreach ($urutan_list as $r){
                if ($r != null && $r != ""){
                    $uuid = $r;
                    DB::table('quiz_sesi_template_report')->where('uuid', $uuid)->update(["urutan"=>$urutan]);
                    $urutan++;
                }
            }
            $respon = array('status'=>true,'message'=>'Urutan Komponen Berhasil Diperbarui!');
            return response()->json($respon);
        }
    }

    
    function get_data_lampiran($uuid){
        $data = DB::table('quiz_sesi_template_lampiran')->where('uuid', $uuid)->first();
        if($data){
            $report = DB::table('quiz_sesi_report')->where('id_report', $data->id_report)->first();
            $respon = array('status'=>true,'data'=>$data, 
                'informasi'=>'Lampiran : '.$report->nama_report);
            return response()->json($respon);
        }
        $respon = array('status'=>false,'message'=>'Data Tidak Ditemukan');
        return response()->json($respon); 
    }

    function submit_insert_lampiran(Request $r){
        if($this->ucc()){
            loadHelper('format');
            $uuid = $this->genUUID(); 
            $id_quiz_template = $r->id_quiz_template;
            $id_report = (int)($r->id_report);
            $tabel_referensi = DB::table('quiz_sesi_report')->where('id_report', $id_report)->first()->tabel_referensi;
            if($tabel_referensi!="-"){
                $exist = DB::table('quiz_sesi_template_lampiran')
                    ->where('id_quiz_template', $id_quiz_template)
                    ->where('id_report',$id_report)
                    ->count();
                if($exist){
                    $respon = array('status'=>false,'message'=>'Lampiran sudah ada!');
                    return response()->json($respon);
                }
            }           
            $jumlah = DB::select(
                "select 
                count(*) as jumlah
                FROM
                    quiz_sesi_template_lampiran AS b,
                    quiz_sesi_report AS a 
                WHERE
                    a.id_report = b.id_report 
                    AND b.id_quiz_template =  $id_quiz_template"
            );
            $urutan = $jumlah[0]->jumlah + 1;
            $record = array( 
                "id_quiz_template"=>$r->id_quiz_template,                                        
                "urutan"=>(int)($urutan),
                "id_report"=>$id_report, 
                "uuid"=>$uuid,
            );
            DB::table('quiz_sesi_template_lampiran')->insert($record);
            $respon = array('status'=>true,'message'=>'Lampiran Berhasil Ditambahkan!');
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }


    function submit_delete_lampiran(Request $r){
        if($this->ucu()){
            //loadHelper('format');
            $id_quiz_template = DB::table('quiz_sesi_template_lampiran')->where('uuid', $r->uuid)->first()->id_quiz_template;
            DB::table('quiz_sesi_template_lampiran')->where('uuid', $r->uuid)->delete();
            $listbaru = DB::table('quiz_sesi_template_lampiran')
                            ->where('id_quiz_template', $id_quiz_template)
                            ->orderby("urutan")->get();
            $urutan = 1;
            foreach ($listbaru as $r){
                $uuid = $r->uuid;
                DB::table('quiz_sesi_template_report')->where('uuid', $uuid)->update(["urutan"=>$urutan]);
                $urutan++;
            }
            $respon = array('status'=>true,'message'=>'Lampiran Berhasil Dihapus!');
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }

    function get_list_lampiran_report($id_quiz_template){
        
        $list = DB::select("select a.nama_report, b.uuid ,
                 b.urutan, a.tabel_referensi 
                        from quiz_sesi_report as a, 
                            quiz_sesi_template_lampiran as b where a.jenis = 2 
                                and a.id_report = b.id_report 
                                    and b.id_quiz_template = $id_quiz_template order by b.urutan asc");
        return view('master-template.drag-lampiran', compact('list'));
    }

    function submit_update_lampiran(Request $r){
        if($this->ucc()){
            $urutan_string =  $r->urutan_list;
            $urutan_list = explode(",", $urutan_string);
            $urutan = 1;
            foreach ($urutan_list as $r){
                if ($r != null && $r != ""){
                    $uuid = $r;
                    DB::table('quiz_sesi_template_lampiran')->where('uuid', $uuid)->update(["urutan"=>$urutan]);
                    $urutan++;
                }
            }
            $respon = array('status'=>true,'message'=>'Urutan Lampiran Berhasil Diperbarui!');
            return response()->json($respon);
        }
    }

    function cek_dummy_template($id){
        ini_set('max_execution_time', '1300');
        $curl = curl_init();
        curl_setopt_array($curl, array(
        CURLOPT_URL => env('GO_API_URL').'/cek-dummy-template/'.$id,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => '',
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 600,
        CURLOPT_CONNECTTIMEOUT=>0,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => 'POST',
            CURLOPT_HTTPHEADER => array(
                'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2OTg1NTc3MDYsImlzcyI6Imh0dHBzOi8vc2ljZXJkYXMud2ViLmlkIiwianRpIjoiZGEyZWYwNTUtODYzNS00OTMyLWJmYTAtNmE0ODRiMTQ4MWU2IiwibmFtZSI6IkFkbWluaXN0cmF0b3IgU0NEIiwic3ViIjoiMjM1MzI2MjM2LTQzNzk0MzA3NTQ4IiwidXNlcm5hbWUiOiJhZG1pbiJ9.Bqb-aApPsbiOkStnt5M10-mc9pM8Ro5YSgDQhiZ5HmYOAogTuc5F9JTHoFhxVcsk2BY3bLkclH2kXoHpMJyPpA'
            ),
        ));
        $response = curl_exec($curl);
        curl_close($curl);
        $data = json_decode($response);
        return response()->json($data);
    }

    function create_dummy_sesi($id){
        ini_set('max_execution_time', '1300');
        $curl = curl_init();
        curl_setopt_array($curl, array(
        CURLOPT_URL => env('GO_API_URL').'/create-dummy-template/'.$id,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => '',
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 600,
        CURLOPT_CONNECTTIMEOUT=>0,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => 'POST',
            CURLOPT_HTTPHEADER => array(
                'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2OTg1NTc3MDYsImlzcyI6Imh0dHBzOi8vc2ljZXJkYXMud2ViLmlkIiwianRpIjoiZGEyZWYwNTUtODYzNS00OTMyLWJmYTAtNmE0ODRiMTQ4MWU2IiwibmFtZSI6IkFkbWluaXN0cmF0b3IgU0NEIiwic3ViIjoiMjM1MzI2MjM2LTQzNzk0MzA3NTQ4IiwidXNlcm5hbWUiOiJhZG1pbiJ9.Bqb-aApPsbiOkStnt5M10-mc9pM8Ro5YSgDQhiZ5HmYOAogTuc5F9JTHoFhxVcsk2BY3bLkclH2kXoHpMJyPpA'
            ),
        ));
        $response = curl_exec($curl);
        curl_close($curl);
        $data = json_decode($response);
        return response()->json($data);
    }

    function export_report_sample(Request $r){
        $id_quiz = $r->id_quiz;
        $id_user = $r->id_user;
        $id_model = $r->id_model;
        ini_set('max_execution_time', '1300');
        // export-report-peserta/{id_quiz}/{id_user}/{id_model}
        $curl = curl_init();
        curl_setopt_array($curl, array(
        CURLOPT_URL => env('GO_API_URL').'/export-report-peserta/'.$id_quiz .'/'.$id_user.'/'.$id_model,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => '',
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 0,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => 'POST',
            CURLOPT_HTTPHEADER => array(
                'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2OTg1NTc3MDYsImlzcyI6Imh0dHBzOi8vc2ljZXJkYXMud2ViLmlkIiwianRpIjoiZGEyZWYwNTUtODYzNS00OTMyLWJmYTAtNmE0ODRiMTQ4MWU2IiwibmFtZSI6IkFkbWluaXN0cmF0b3IgU0NEIiwic3ViIjoiMjM1MzI2MjM2LTQzNzk0MzA3NTQ4IiwidXNlcm5hbWUiOiJhZG1pbiJ9.Bqb-aApPsbiOkStnt5M10-mc9pM8Ro5YSgDQhiZ5HmYOAogTuc5F9JTHoFhxVcsk2BY3bLkclH2kXoHpMJyPpA'
            ),
        ));
        $response = curl_exec($curl);
        curl_close($curl);
        $data = json_decode($response);
        return response()->json($data);
    }

    function generate_preview_template_pdf($uuid, $model){
        $url = env('GO_API_URL').'/preview-report-template/'.$uuid.'/'.$model;
        $path = env('PATH_REPORT_BARU'); 
        $filename = $uuid.'.pdf';

        $command = env('WKHTML').' --footer-spacing 3 -L 10 -R 10 -T 20  -B 20 --footer-left "Si Cerdas Indonesia"  --footer-font-size 9 --footer-center [page]/[topage] -O Portrait  -s A4 '.$url.' '.$path.$filename;
        echo $command;
        $process = new Process($command);
        $process->run();
    }
}
