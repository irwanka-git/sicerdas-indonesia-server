<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;

class SoalSSCTRemajaController   extends Controller
{
     
    //######################### SOAL TEST KOGNITIF #####################################
    function index(){
    	$pagetitle = "Soal SSCT Remaja";
    	$smalltitle = "Pengaturan Soal";
    	return view('manajeman-soal.index-ssct-remaja', compact('pagetitle','smalltitle'));
    }

    function datatable(Request $r){
        
        $filter = "";
        if (request()->has('search')) {
            $search = request('search');
            $keyword = $search['value'];
            if(strlen($keyword)>=3){
                $keyword = strtolower($keyword);
                $filter = " and ( lower(a.subjek_penilaian) like '%$keyword%'   
                             or   lower(a.komponen) like '%$keyword%'                     
                            ) ";
            }   
        }

         $sql_union = "select a.* from soal_ssct_remaja as a where a.id_soal > 0 $filter  ";
         $query = DB::table(DB::raw("($sql_union) as x order by urutan"))
                    ->select([
                        'id_soal',
                        'urutan',
                        'komponen',
                        'subjek_penilaian',
                        'aspek',
                        'sikap_negatif1',
                        'sikap_negatif2',
                        'sikap_negatif3',
                        'sikap_positif1',
                        'sikap_positif2',
                        'sikap_positif3',
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
            ->editColumn('subjek_penilaian', function($q){
                return $q->subjek_penilaian;
            })
            ->editColumn('sikap_negatif', function($q){
                $list ="<ul>";
                $list .= "<li>".$q->sikap_negatif1."</li>";
                $list .= "<li>".$q->sikap_negatif2."</li>";
                $list .= "<li>".$q->sikap_negatif3."</li>";
                $list .= "</ul>";
                return $list;
            })
            ->editColumn('sikap_positif', function($q){
                $list ="<ul>";
                $list .= "<li>".$q->sikap_positif1."</li>";
                $list .= "<li>".$q->sikap_positif2."</li>";
                $list .= "<li>".$q->sikap_positif3."</li>";
                $list .= "</ul>";
                return $list;
            })
            ->editColumn('subjek_penilaian', function($q){
                return $q->subjek_penilaian;
            })
            ->addIndexColumn()
            ->rawColumns(['action','subjek_penilaian', 'sikap_negatif', 'sikap_positif'])
            ->make(true);
    } 

    function get_data($uuid){
    	$data = DB::table('soal_ssct_remaja')->where('uuid', $uuid)->first();
        if($data){
            $respon = array('status'=>true,'data'=>$data, 
            	'informasi'=>'Informasi Soal : Subjek Penilaian Ke '.$data->urutan);
            return response()->json($respon);
        }
        $respon = array('status'=>false,'message'=>'Data Tidak Ditemukan');
        return response()->json($respon);
    }

    function submit_insert(Request $r){
    	if($this->ucc()){
	    	loadHelper('format');
	    	$uuid = $this->genUUID();
            $subjek_penilaian = str_replace("<p>", "", $r->subjek_penilaian);
            $subjek_penilaian = str_replace("</p>", "", $subjek_penilaian);
            //return $subjek_penilaian;
	    	$record = array(
                "urutan"=>(int)$r->urutan, 
                "subjek_penilaian"=>trim($subjek_penilaian), 
                "komponen"=>trim($r->komponen), 
                "aspek"=>strtoupper($r->aspek), 
                "sikap_negatif1"=>trim($r->sikap_negatif1), 
                "sikap_positif1"=>trim($r->sikap_positif1),
                "sikap_negatif2"=>trim($r->sikap_negatif2), 
                "sikap_positif2"=>trim($r->sikap_positif2),
                "sikap_negatif3"=>trim($r->sikap_negatif3), 
                "sikap_positif3"=>trim($r->sikap_positif3), 
	    		"uuid"=>$uuid);

	    	DB::table('soal_ssct_remaja')->insert($record);
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
            $subjek_penilaian = str_replace("<p>", "", $r->subjek_penilaian);
            $subjek_penilaian = str_replace("</p>", "", $subjek_penilaian);
            //return $subjek_penilaian;
	    	$record = array(
               "urutan"=> $r->urutan, 
                "subjek_penilaian"=>trim($subjek_penilaian), 
                "komponen"=>trim($r->komponen), 
                "aspek"=>strtoupper($r->aspek),  
                "sikap_negatif1"=>trim($r->sikap_negatif1), 
                "sikap_positif1"=>trim($r->sikap_positif1),
                "sikap_negatif2"=>trim($r->sikap_negatif2), 
                "sikap_positif2"=>trim($r->sikap_positif2),
                "sikap_negatif3"=>trim($r->sikap_negatif3), 
                "sikap_positif3"=>trim($r->sikap_positif3), 
            );

	    	DB::table('soal_ssct_remaja')->where('uuid', $uuid)->update($record);
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
            DB::table('soal_ssct_remaja')->where('uuid', $uuid)->delete();
            $respon = array('status'=>true,'message'=>'Data Berhasil Dihapus!');          
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }
    //#######################################################################################

}
