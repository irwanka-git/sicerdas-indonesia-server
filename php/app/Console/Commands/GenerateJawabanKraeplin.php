<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use DB;

class GenerateJawabanKraeplin extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'generate:kraeplin';

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
 
    
    public function handle()
    {
        $sql = "select * from soal_kraeplin order by urutan";
        $data = DB::select($sql);

        foreach($data as $r){
        	$pertanyaan = $r->pertanyaan;
        	echo  $pertanyaan."\n";
        	$jawaban =  $this->generate_jawaban($pertanyaan);
        	$uuid = $r->uuid;
        	echo $jawaban ."\n";
        	echo $uuid ."\n";
        	DB::table('soal_kraeplin')->where('uuid',trim($uuid))->update(['pilihan_jawaban'=> trim($jawaban)]);
        }
    }
 	
 	public function generate_jawaban($pertanyaan){
 		$len = strlen($pertanyaan);
 		$arr = [];
 		for($i=0; $i<=$len - 1; $i++){
 			array_push($arr, substr($pertanyaan,$i,1));
 		}
 		$arr_jawaban = [];
 		$angka_awal = $arr[0];
 		$jawaban_string = "";
 		for($i=1; $i<=$len - 1; $i++){
 			$jawaban = ($angka_awal + $arr[$i])  % 10;
 			array_push($arr_jawaban,$jawaban);
 			$angka_awal = $arr[$i];
 			$jawaban_string = $jawaban_string . $jawaban;
 		}
 		return $jawaban_string;
 	}

}
