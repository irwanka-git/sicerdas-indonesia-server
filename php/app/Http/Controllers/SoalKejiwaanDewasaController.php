<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;

class SoalKejiwaanDewasaController   extends Controller
{
     
    //######################### SOAL TEST KOGNITIF #####################################
    function index(){
    	$pagetitle = "Soal Tes Kejiwaan Dewasa Indonesia";
    	$smalltitle = "<p>Pilihan Jawaban : (1) Tidak  (2) Kadang (3) Ya</p>";
    	return view('manajeman-soal.index-soal-kejiwaan-dewasa', compact('pagetitle','smalltitle'));
    }

    function datatable(Request $r){
        
        $filter = "";
        if (request()->has('search')) {
            $search = request('search');
            $keyword = $search['value'];
            if(strlen($keyword)>=3){
                $keyword = strtolower($keyword);
                $filter = " and ( lower(a.unsur) like '%$keyword%' or lower(b.nama) like '%$keyword%') ";
            }   
        }

         $sql_union = "select a.id_soal, a.id_model, a.urutan, a.unsur , b.nama as model                 
         from soal_kejiwaan_dewasa as a, ref_model_kejiwaan_dewasa as b where a.id_model = b.id  $filter  ";
         $query = DB::table(DB::raw("($sql_union) as x order by id_model, urutan"))
                    ->select([
                        'id_soal',
                        'id_model',
                        'urutan',
                        'model',
                        'unsur', 
                    ]);

         return Datatables::of($query)
         ->editColumn('model', function($q){
            return $q->model. " (Model ".$q->id_model.")";
         })
            ->addIndexColumn() 
            ->make(true);
    }

    // <a href="#"><i class="align-middle" data-feather="edit-2"></i></a>
    //                                             <a href="#"><i class="align-middle" data-feather="trash"></i></a>
}
