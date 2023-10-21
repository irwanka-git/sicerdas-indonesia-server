<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use DB;

class SkoringJob extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'skoring:job';

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

    /**
     * Execute the console command.
     
        php artisan skoring:peminatan_sma
        php artisan skoring:peminatan_man
        php artisan skoring:peminatan_smk
        php artisan skoring:penjurusan_kuliah
        php artisan skoring:peminatan_lengkap
        php artisan skoring:peminatan_sma_v2
        php artisan skoring:penjurusan_kuliah_v2
        php artisan skoring:peminatan_smk_v2
        php artisan skoring:test_iq_dan_eq
        php artisan skoring:seleksi_karyawan
        php artisan skoring:clear_table_jawaban

     * @return mixed
     */
     

    public function handle()
    {   
         $status = DB::table('status_cronjob')->first()->status;
         if($status==0){
            echo "SKORING STARTING........\n";
            echo "*********************************************************\n";
            echo "*********************************************************\n";
            $status = DB::table('status_cronjob')->update(['status'=>1]);

             $this->call('skoring:peminatan_sma');
             $this->call('skoring:peminatan_man');
             $this->call('skoring:peminatan_smk');
             $this->call('skoring:penjurusan_kuliah');
             $this->call('skoring:peminatan_lengkap');
             $this->call('skoring:peminatan_sma_v2');
             $this->call('skoring:penjurusan_kuliah_v2');
             $this->call('skoring:peminatan_smk_v2');
             $this->call('skoring:test_iq_dan_eq');
             $this->call('skoring:seleksi_karyawan');
             $this->call('skoring:peminatan_sma_v3');
             $this->call('skoring:penjurusan_kuliah_v3'); 
             $status = DB::table('status_cronjob')->update(['status'=>0]);
             $this->call('skoring:clear_table_jawaban');
             


             echo "\n*********************************************************\n";
             echo "*********************************************************\n";
             echo "SKORING COMPELETED........\n";
         }else{
            echo "Another Skoring is Running!";
         }         
    }

    
}
