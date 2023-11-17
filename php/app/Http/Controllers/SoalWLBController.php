<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use Session;
use Datatables;
use Crypt;

class SoalWLBController   extends Controller
{
     
    //######################### SOAL TEST KOGNITIF #####################################
    function index(){
    	$pagetitle = "Soal Work Life Balanced";
    	$smalltitle = "<p>Pilihan Jawaban : (1) Tidak  (2) Kadang (3) Ya</p>";
    	return view('manajeman-soal.index-soal-wlb', compact('pagetitle','smalltitle'));
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
         from soal_wlb as a, ref_model_wlb as b where a.id_model = b.id  $filter  ";
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
            return "<small>(Keseimbangan ".$q->id_model.")</small><br>". $q->model;
         })
            ->addIndexColumn() 
            ->rawColumns(['action','model'])
            ->make(true);
    }

    // <a href="#"><i class="align-middle" data-feather="edit-2"></i></a>
    //                                             <a href="#"><i class="align-middle" data-feather="trash"></i></a>
}
