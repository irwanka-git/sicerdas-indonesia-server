<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;
use Auth;
class ProfilBiroController extends Controller
{
     
    //######################### PETUNJUK SOAL #####################################
    function index(){
    	$pagetitle = "Pengaturan Profil Biro";
    	$smalltitle = "Pengaturan Petunjuk Soal";
    	return view('manajemen-user.profil-biro', compact('pagetitle','smalltitle'));
    }

    function profil_biro($uuid){
    	$pagetitle = "Pengaturan Profil Biro";
    	$smalltitle = "Pengaturan Petunjuk Soal";
    	return view('manajemen-user.profil-biro-uuid', compact('pagetitle','smalltitle','uuid'));
    }
    
    function submit_update(Request $r){
    	if($this->ucu()){
	    	$record = array(
	    		"nama_pengguna"=>trim($r->nama_pengguna), 
	    		"alamat"=>trim($r->alamat), 
	    		"email"=>trim($r->email), 
	    		"telp"=>trim($r->telp), 
	    		"kop_biro"=>trim($r->kop_biro), 
	    		"cover_biro"=>trim($r->cover_biro), 
	    	);

	    	DB::table('users')->where('id', Auth::user()->id)->update($record);
	    	$respon = array('status'=>true,'message'=>'Data User Berhasil Disimpan!', 
	    		'_token'=>csrf_token());
        	return response()->json($respon);
    	}else{
    		$respon = array('status'=>false,'message'=>'Akses Ditolak!');
        	return response()->json($respon);
    	}
    }

    function submit_update_profil(Request $r){
    	$uuid = $r->uuid;
    	if($this->ucu()){
	    	$record = array(
	    		"nama_pengguna"=>trim($r->nama_pengguna), 
	    		"alamat"=>trim($r->alamat), 
	    		"email"=>trim($r->email), 
	    		"telp"=>trim($r->telp), 
	    		"kop_biro"=>trim($r->kop_biro), 
	    		"cover_biro"=>trim($r->cover_biro), 
	    		"cover_biro_gambar"=>trim($r->cover_biro_gambar), 
	    	);

	    	DB::table('users')->where('uuid', $uuid)->update($record);
	    	$respon = array('status'=>true,'message'=>'Data Profil Biro Berhasil Disimpan!', 
	    		'_token'=>csrf_token());
        	return response()->json($respon);
    	}else{
    		$respon = array('status'=>false,'message'=>'Akses Ditolak!');
        	return response()->json($respon);
    	}
    }
 

}
