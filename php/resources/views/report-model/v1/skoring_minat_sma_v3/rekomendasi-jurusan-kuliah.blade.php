<table width="100%" cellpadding="0" cellspacing="5">
    <thead>
        <tr class="box-header-010">
            <th colspan="2" class="text-center box-header-010 ">REKOMENDASI JURUSAN KULIAH</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td width="50%" class="box-cell-info-010" align="center"><b>KELOMPOK ILMU SAINS DAN TEKNOLOGI</b></td>
            <td width="50%" class="box-cell-info-010" align="center"><b>KELOMPOK ILMU SOSIAL HUMANIORA</b></td>
        </tr>
        <tr>
            <td class="box-cell-info-010">
                <?php
                    $list_ipa = DB::select("SELECT *  FROM soal_minat_kuliah_eksakta");
                    $data_jawaban = $data_sesi->jawaban_skoring;
                    $obj = json_decode($data_jawaban);
                    $jawaban_minat = $obj->SKALA_PMK_MINAT_ILMU_ALAM;
                ?>
                <ol>
                @foreach ($jawaban_minat as $key => $value)
                    @foreach ($list_ipa as $l)
                        @if($l->urutan ==(int)$value->jawaban)
                        <li>{{$l->jurusan}}</li>
                        @endif
                    @endforeach
                @endforeach
                </ol>
            </td>
            <td class="box-cell-info-010" >
                <?php
                    $list_ips = DB::select("SELECT *  FROM soal_minat_kuliah_sosial");
                    $data_jawaban = $data_sesi->jawaban_skoring;
                    $obj = json_decode($data_jawaban);
                    $jawaban_minat = $obj->SKALA_PMK_MINAT_ILMU_SOSIAL;
                ?>
                <ol>
                @foreach ($jawaban_minat as $key => $value)
                    @foreach ($list_ips as $l)
                        @if($l->urutan ==(int)$value->jawaban)
                        <li>{{$l->jurusan}}</li>
                        @endif
                    @endforeach
                @endforeach
                </ol>
            </td>
         </tr>
         <tr>
             <td class="box-cell-info-010"  colspan="2">
                 <b>SEKOLAH KEDINASAN : </b><br>
                 <?php
                    $jumlah_pilihan = 3;
                    $list_sekolah = DB::select("SELECT *  FROM ref_sekolah_dinas");
                    $sekolah = [];
                    foreach($list_sekolah as $s){
                        $sekolah[$s->no]['akronim'] = $s->akronim;
                        $sekolah[$s->no]['nama'] = $s->nama_sekolah_dinas;
                    }
                ?>
                <ol>
                <?php 
                for ($i=1;$i<=$jumlah_pilihan;$i++){
                    $field_name = 'minat_dinas'.$i;
                    $nomor = $data_skoring->$field_name;
                ?>
                    <li> {{$sekolah[$nomor]['akronim']}} <small> - {{$sekolah[$nomor]['nama']}}</small></li>
                <?php
                }
                ?>
                </ol>
             </td>
         </tr>
         <tr>
             <td class="box-cell-info-010"  colspan="2">
                Lebih disarankan ke Mata Pelajaran<br>
                 <b>{{str_replace("_"," ",$data_skoring->rekom_akhir)}}</b>
             </td>
         </tr>
    </tbody>
</table>
