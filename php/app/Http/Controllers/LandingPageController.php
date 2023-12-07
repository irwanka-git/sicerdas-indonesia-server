<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;
use Auth;
use Image;

class LandingPageController extends Controller
{
    function index(){
    	 $pagetitle = "Beranda";
    	 return view('landing.index', compact('pagetitle'));
    }

    function detil_paket($uuid){
        $paket = DB::table('tarif_paket')->where('uuid', $uuid)->first();
        $sesi = DB::table('tarif_paket_rinci')->where('id_tarif', $paket->id_tarif)->get();
        return view('landing.detil', compact('paket','sesi'));
    }
        
}
 