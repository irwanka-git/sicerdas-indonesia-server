<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;

class SoalModeBelajarController  extends Controller
{
     
    //######################### SOAL MODE BELAJAR#####################################
    function index(){
    	$pagetitle = "Soal Mode Belajar";
    	$smalltitle = "Pengaturan Soal Mode Belajar";
    	return view('manajeman-soal.index-soal-mode-belajar', compact('pagetitle','smalltitle'));
    }

    function datatable(Request $r){
        
        $filter = "";
        if (request()->has('search')) {
            $search = request('search');
            $keyword = $search['value'];
            if(strlen($keyword)>=3){
                $keyword = strtolower($keyword);
                $filter = " and ( lower(deskripsi) like '%$keyword%') ";
            }   
        }

       
         $sql_union = "select * from soal_mode_belajar where urutan > 0 $filter order by urutan";

         $query = DB::table(DB::raw("($sql_union) as x"))
                    ->select([
                        'urutan',
                        'soal',
                        'deskripsi',
                        'suasana_belajar',
                        'pilihan_a',
                        'pilihan_b',
                        'pilihan_c',
                        'pilihan_d',
                        'pilihan_e',
                        'uuid',
                    ]);

         return Datatables::of($query)
            ->addColumn('action', function ($query) {
                    $edit = ""; $delete = "";
                    $action = '';
                    if($this->ucu()){
                        $edit = '<button data-bs-toggle="modal" data-uuid="'.$query->uuid.'" data-bs-target="#modal-edit" class="btn btn-outline-secondary btn-outline btn-sm" type="button"><i class="las la-pen"></i></button>';
                    }
                    if($this->ucd()){
                        $delete = '<button  data-uuid="'.$query->uuid.'" class="btn btn-outline-secondary btn-sm btn-konfirm-delete" type="button"><i class="las la-trash"></i></button>';
                    }
                    $action =  $action." ".$edit." ".$delete;
                    if ($action==""){$action='<a href="#" class="act"><i class="la la-lock"></i></a>'; }
                        return $action;
            })
            ->editColumn('soal', function ($q) {
                return $q->soal."<br>Aspek: <b>".$q->suasana_belajar."</b>";
             })
            ->addColumn('pilihan_jawaban', function ($q) {
                $pilihan_a = explode(":", $q->pilihan_a);
                 
                $_pilihan_a = "";
                if (is_array($pilihan_a)){
                    if (count($pilihan_a)==2){
                        $_pilihan_a = $pilihan_a[0];
                    } 
                }
                

                $pilihan_b = explode(":", $q->pilihan_b);
                $_pilihan_b = "";
                if (is_array($pilihan_b)){
                    if (count($pilihan_b)==2){
                        $_pilihan_b = $pilihan_b[0];
                    } 
                }

                $pilihan_c = explode(":", $q->pilihan_c);
                $_pilihan_c = "";
                if (is_array($pilihan_c)){
                    if (count($pilihan_c)==2){
                        $_pilihan_c = $pilihan_c[0];
                    } 
                }

                $pilihan_d = explode(":", $q->pilihan_d);
                $_pilihan_d = "";
                if (is_array($pilihan_d)){
                    if (count($pilihan_d)==2){
                        $_pilihan_d = $pilihan_d[0];
                    } 
                }

                $pilihan_e = explode(":", $q->pilihan_e);
                $_pilihan_e = "";
                if (is_array($pilihan_d)){
                    if (count($pilihan_e)==2){
                        $_pilihan_e = $pilihan_e[0];
                    } 
                }
               
                $list = "<li>A. ".trim($_pilihan_a)."</li>";
                $list .= "<li>B. ".trim($_pilihan_b)."</li>";
                $list .= "<li>C. ".trim($_pilihan_c)."</li>";
                $list .= "<li>D. ".trim($_pilihan_d)."</li>";
                $list .= "<li>E. ".trim($_pilihan_e)."</li>";

                $res = "<ul>".$list."</ul>";
                return $res;
             })
             ->addColumn('nama_prioritas', function ($q) {
                $pilihan_a = explode(":", $q->pilihan_a);
                $_pilihan_a = "";
                if (count($pilihan_a)==2){
                    $_pilihan_a = $pilihan_a[1];
                } 

                $pilihan_b = explode(":", $q->pilihan_b);
                $_pilihan_b = "";
                if (count($pilihan_b)==2){
                    $_pilihan_b = $pilihan_b[1];
                } 

                $pilihan_c = explode(":", $q->pilihan_c);
                $_pilihan_c = "";
                if (count($pilihan_c)==2){
                    $_pilihan_c = $pilihan_c[1];
                } 

                $pilihan_d = explode(":", $q->pilihan_d);
                $_pilihan_d = "";
                if (count($pilihan_d)==2){
                    $_pilihan_d = $pilihan_d[1];
                } 

                $pilihan_e = explode(":", $q->pilihan_e);
                $_pilihan_e = "";
                if (count($pilihan_e)==2){
                    $_pilihan_e = $pilihan_e[1];
                } 

                $list = "<li>A. ".trim($_pilihan_a)."</li>";
                $list .= "<li>B. ".trim($_pilihan_b)."</li>";
                $list .= "<li>C. ".trim($_pilihan_c)."</li>";
                $list .= "<li>D. ".trim($_pilihan_d)."</li>";
                $list .= "<li>E. ".trim($_pilihan_e)."</li>";

                $res = "<ul>".$list."</ul>";
                return $res;
             })
            ->addIndexColumn()
            ->rawColumns(['action','deskripsi', 'soal', 'pilihan_jawaban','nama_prioritas'])
            ->make(true);
    }

 
 
    function get_data($uuid){
    	$data = DB::table('soal_mode_belajar')->where('uuid', $uuid)->first();
        if($data){
            $respon = array('status'=>true,'data'=>$data, 
            	'informasi'=>'Informasi Soal : Soal Ke '.$data->urutan);
            return response()->json($respon);
        }
        $respon = array('status'=>false,'message'=>'Data Tidak Ditemukan');
        return response()->json($respon);
    }

    function submit_insert(Request $r){
    	if($this->ucc()){
	    	loadHelper('format');
	    	$uuid = $this->genUUID();
            $deskripsi = str_replace("<p>", "", $r->deskripsi);
            $deskripsi = str_replace("</p>", "", $deskripsi);
            
            $soal = trim($r->soal);
            $suasana_belajar = trim($r->suasana_belajar);

            $pilihan_a = trim($r->pilihan_a);
            $pilihan_b = trim($r->pilihan_b);
            $pilihan_c = trim($r->pilihan_c);
            $pilihan_d = trim($r->pilihan_d);
            $pilihan_e = trim($r->pilihan_e);

            $urutan = (int)($r->urutan);

           
            //return $pernyataan;
	    	$record = array(
                "urutan"=>$urutan,            
                "soal"=>$soal,            
                "deskripsi"=>$deskripsi,  
                "suasana_belajar"=>$suasana_belajar,          
                "urutan"=>$urutan,            
                "pilihan_a"=>$pilihan_a,            
                "pilihan_b"=>$pilihan_b,            
                "pilihan_c"=>$pilihan_c,            
                "pilihan_d"=>$pilihan_d,            
                "pilihan_e"=>$pilihan_e,  

	    		"uuid"=>$uuid);

	    	DB::table('soal_mode_belajar')->insert($record);
	    	$respon = array('status'=>true,'message'=>'Soal Berhasil Ditambahkan!');
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

            $deskripsi = str_replace("<p>", "", $r->deskripsi);
            $deskripsi = str_replace("</p>", "", $deskripsi);
            
            $soal = trim($r->soal);
            $suasana_belajar = trim($r->suasana_belajar);

            $pilihan_a = trim($r->pilihan_a);
            $pilihan_b = trim($r->pilihan_b);
            $pilihan_c = trim($r->pilihan_c);
            $pilihan_d = trim($r->pilihan_d);
            $pilihan_e = trim($r->pilihan_e);

            $urutan = (int)($r->urutan);
            
            //return $pernyataan;
	    	$record = array(
                "urutan"=>$urutan,            
                "soal"=>$soal,            
                "deskripsi"=>$deskripsi,  
                "suasana_belajar"=>$suasana_belajar,          
                "urutan"=>$urutan,            
                "pilihan_a"=>$pilihan_a,            
                "pilihan_b"=>$pilihan_b,            
                "pilihan_c"=>$pilihan_c,            
                "pilihan_d"=>$pilihan_d,            
                "pilihan_e"=>$pilihan_e, 
            );

	    	DB::table('soal_mode_belajar')->where('uuid', $uuid)->update($record);
	    	$respon = array('status'=>true,'message'=>'Soal Berhasil Disimpan!', 
	    		'_token'=>csrf_token());
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
            DB::table('soal_mode_belajar')->where('uuid', $uuid)->delete();
            $respon = array('status'=>true,'message'=>'Data Berhasil Dihapus!');          
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }
    //#######################################################################################

}
