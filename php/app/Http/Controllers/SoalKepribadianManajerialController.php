<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;

class SoalKepribadianManajerialController  extends Controller
{
     
    //######################### SOAL TEST KOGNITIF #####################################
    function index(){
    	$pagetitle = "Soal Kepribadian Manajerial";
    	$smalltitle = "Pengaturan Soal Kepribadian Manajerial";
    	return view('manajeman-soal.index-kepribadian-manajerial', compact('pagetitle','smalltitle'));
    }

    function datatable(Request $r){
        
        $filter = "";
        if (request()->has('search')) {
            $search = request('search');
            $keyword = $search['value'];
            if(strlen($keyword)>=3){
                $keyword = strtolower($keyword);
                $filter = " and ( lower(a.pernyataan) like '%$keyword%') ";
            }   
        }

       
         $sql_union = "select a.id_soal , a.id_komponen, a.urutan, a.pernyataan,a.uuid, b.komponen as nama_komponen  
         from soal_kepribadian_manajerial as a, ref_komponen_kepribadian_manajerial as b where a.id_komponen = b.id and a.id_soal > 0  $filter order by a.id_soal";

         $query = DB::table(DB::raw("($sql_union) as x"))
                    ->select([
                        'id_soal',
                        'nama_komponen',
                        'urutan',
                        'pernyataan',
                        'uuid',
                    ]);

         return Datatables::of($query)
            ->addColumn('action', function ($query) {
                    $edit = ""; $delete = "";
                    $action = ''; //'<button data-bs-toggle="modal" data-uuid="'.$query->uuid.'" data-bs-target="#modal-lihat-soal" class="btn btn-outline-secondary btn-outline btn-sm" type="button"><i class="las la-eye"></i></button>';
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
            ->editColumn('pernyataan', function ($q){
                   $pernyataan =   "(".$q->urutan.")  ".$q->pernyataan;
                    return $pernyataan;
            })
            ->addIndexColumn()
            ->rawColumns(['action','pernyataan'])
            ->make(true);
    }

 

    function view_soal($uuid){
        // $soal = DB::table('soal_kepribadian_manajerial')->where('uuid', $uuid)->first();
        $soal=DB::select("select a.*, b.nama_komponen  from soal_kepribadian_manajerial as a, ref_komponen_karakteristik_pribadi as b where a.id_komponen = b.id_komponen and a.uuid ='$uuid' ");
        if(count($soal)==1){
            $soal=$soal[0];

            return view('manajeman-soal.view-karakter-pribadi',compact('soal'));
        }
        return '<center><b>Maaf, Soal Tidak Ditemukan</b></center>';
    }

    function get_data($uuid){
    	$data = DB::table('soal_kepribadian_manajerial')->where('uuid', $uuid)->first();
        if($data){
            $respon = array('status'=>true,'data'=>$data, 
            	'informasi'=>'Informasi Soal : Pernyataan Ke '.$data->urutan);
            return response()->json($respon);
        }
        $respon = array('status'=>false,'message'=>'Data Tidak Ditemukan');
        return response()->json($respon);
    }

    function submit_insert(Request $r){
    	if($this->ucc()){
	    	loadHelper('format');
	    	$uuid = $this->genUUID();
            $pernyataan = str_replace("<p>", "", $r->pernyataan);
            $pernyataan = str_replace("</p>", "", $pernyataan);
           
            //return $pernyataan;
	    	$record = array(
                "id_komponen"=>(int)$r->id_komponen,
                "urutan"=>(int)$r->urutan, 
                "pernyataan"=>trim($pernyataan),               
	    		"uuid"=>$uuid);

	    	DB::table('soal_kepribadian_manajerial')->insert($record);
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
            $pernyataan = str_replace("<p>", "", $r->pernyataan);
            $pernyataan = str_replace("</p>", "", $pernyataan);
            //return $pernyataan;
	    	$record = array(
                "id_komponen"=>(int)$r->id_komponen,
                "urutan"=>(int)$r->urutan, 
                "pernyataan"=>trim($pernyataan),   
            );

	    	DB::table('soal_kepribadian_manajerial')->where('uuid', $uuid)->update($record);
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
            DB::table('soal_kepribadian_manajerial')->where('uuid', $uuid)->delete();
            $respon = array('status'=>true,'message'=>'Data Berhasil Dihapus!');          
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }
    //#######################################################################################

}
