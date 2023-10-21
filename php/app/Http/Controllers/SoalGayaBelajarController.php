<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;

class SoalGayaBelajarController  extends Controller
{
     
    //######################### SOAL TEST KOGNITIF #####################################
    function index(){
    	$pagetitle = "Soal Gaya Belajar";
    	$smalltitle = "Pengaturan Soal Gaya Belajar";
    	return view('manajeman-soal.index-soal-gaya-belajar', compact('pagetitle','smalltitle'));
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

       
         $sql_union = "select  
                a.urutan, 
                a.pernyataan,
                a.pilihan_a,
                a.pilihan_b,
                a.pilihan_c,
                a.uuid
                from soal_gaya_belajar as a 
                where a.id_soal > 0 
                $filter order by a.urutan";

         $query = DB::table(DB::raw("($sql_union) as x"))
                    ->select([
                        'urutan',
                        'pernyataan',
                        'pilihan_a',
                        'pilihan_b',
                        'pilihan_c',
                        'uuid',
                    ]);

         return Datatables::of($query)
            ->addColumn('action', function ($query) {
                    $edit = ""; $delete = "";
                    $action = "";
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
                    return $q->pernyataan;
            })
            ->editColumn('pilihan', function ($q){
                   $pilihan =   "<li>".$q->pilihan_a."</li>";
                   $pilihan .=   "<li>".$q->pilihan_b."</li>";
                   $pilihan .=   "<li>".$q->pilihan_c."</li>";
                    return "<ul>".$pilihan."</ul>";
            })
            ->addIndexColumn()
            ->rawColumns(['action','pernyataan', 'pilihan'])
            ->make(true);
    }

 

    function view_soal($uuid){
        // $soal = DB::table('soal_gaya_belajar')->where('uuid', $uuid)->first();
        $soal=DB::select("select a.*, b.nama_komponen  from soal_gaya_belajar as a, ref_komponen_karakteristik_pribadi as b where a.id_komponen = b.id_komponen and a.uuid ='$uuid' ");
        if(count($soal)==1){
            $soal=$soal[0];

            return view('manajeman-soal.view-karakter-pribadi',compact('soal'));
        }
        return '<center><b>Maaf, Soal Tidak Ditemukan</b></center>';
    }

    function get_data($uuid){
    	$data = DB::table('soal_gaya_belajar')->where('uuid', $uuid)->first();
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
            $pilihan_a = trim($r->pilihan_a);
            $pilihan_b = trim($r->pilihan_b);
            $pilihan_c = trim($r->pilihan_c);
           
            //return $pernyataan;
	    	$record = array(
                "urutan"=>(int)$r->urutan, 
                "pernyataan"=>trim($pernyataan),               
                "pilihan_a"=>trim($pilihan_a),               
                "pilihan_b"=>trim($pilihan_b),               
                "pilihan_c"=>trim($pilihan_c),               
	    		"uuid"=>$uuid);

	    	DB::table('soal_gaya_belajar')->insert($record);
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
            $pilihan_a = trim($r->pilihan_a);
            $pilihan_b = trim($r->pilihan_b);
            $pilihan_c = trim($r->pilihan_c);
            //return $pernyataan;
	    	$record = array( 
                "urutan"=>(int)$r->urutan, 
                "pernyataan"=>trim($pernyataan),  
                "pilihan_a"=>trim($pilihan_a),               
                "pilihan_b"=>trim($pilihan_b),               
                "pilihan_c"=>trim($pilihan_c) 
            );

	    	DB::table('soal_gaya_belajar')->where('uuid', $uuid)->update($record);
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
            DB::table('soal_gaya_belajar')->where('uuid', $uuid)->delete();
            $respon = array('status'=>true,'message'=>'Data Berhasil Dihapus!');          
            return response()->json($respon);
        }else{
            $respon = array('status'=>false,'message'=>'Akses Ditolak!');
            return response()->json($respon);
        }
    }
    //#######################################################################################

}
