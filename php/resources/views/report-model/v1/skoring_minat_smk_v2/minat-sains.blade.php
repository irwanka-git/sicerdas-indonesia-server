<?php 
loadHelper('skoring,function');
?>
<div  style="padding-left:5px;padding-right:5px;">
    <div class="box-title-005">
        <b>PEMINATAN JURUSAN ILMU SAINS TEKNOLOGI</b>
    </div> 
    <br>
    <div class="box-subtitle-005">
        Keterangan: Urutan nomor 1 (satu) adalah jurusan kuliah yang paling dominan, urutan momor 2 (dua) adalah jurusan kuliah kedua yang dominan. Begitupun yang seterusnya 
    </div>
</div>
<table width="100%" cellpadding="0" cellspacing="5">
    <tbody>
        <?php
        $list = DB::select("SELECT *  FROM soal_minat_kuliah_eksakta");
        $data_jawaban = $data_sesi->jawaban_skoring;
        $obj = json_decode($data_jawaban);
        $jawaban_minat = $obj->SKALA_PMK_MINAT_ILMU_ALAM;
        ?>
        @foreach ($jawaban_minat as $key => $value)
            <tr>
                <td width="5%" class="box-urutan-005">
                    {{$value->urutan}}
                </td>
                <td width="10%" class="box-cell-icon2 box-icon-005" align="center" valign="align-middle">
                    @foreach ($list as $l)
                        @if($l->urutan ==(int)$value->jawaban)
                        <img src="{{asset('gambar/'.$l->gambar)}}" height="60px">
                        @endif
                    @endforeach
                </td>
                <td width="85%" class="box-cell-info">
                    @foreach ($list as $l)
                        @if($l->urutan ==(int)$value->jawaban)
                        <b>{{$l->minat}}</b> <small> - {{$l->deskripsi_minat}}</small>
                        @endif
                    @endforeach
                </td>
            </tr>
         @endforeach
    </tbody>
</table>
 