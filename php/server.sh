#!/bin/bash
cd /home/scd-id
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