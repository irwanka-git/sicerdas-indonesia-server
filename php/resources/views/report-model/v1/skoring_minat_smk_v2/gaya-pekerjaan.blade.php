<?php 
loadHelper('skoring,function');
$id_user = $data_skoring->id_user;
$id_quiz = $data_skoring->id_quiz;

$list_klasifikasi= DB::table('ref_klasifikasi_gaya_kerja')->select('akronim')->get();
$gaya_pekerjaan = DB::select("SELECT A
    .ID,
    A.rangking,
    C.nama_komponen,
    C.kode,
    C.cetak_komponen,
    A.skor,
    A.klasifikasi 
FROM
    ref_skoring_gaya_pekerjaan AS A,
    ref_komponen_gaya_pekerjaan AS C 
WHERE
    C.kode = A.kode 
    AND A.id_quiz = $id_quiz 
    AND A.id_user = $id_user 
ORDER BY
    A.rangking");

?>
<div  style="padding-left:5px;padding-right:5px;">
    <div class="box-title-004">
        <b>POLA GAYA PEKERJAAN</b>
    </div> 
    <br>
    <div class="box-subtitle-004">
        <b>Semakin tinggi skor gaya pekerjaan, semakin tinggi pula urutan pola gaya pekerjaan (3 gaya pertama adalah yang dominan)</b>
    </div>
</div>
<table width="100%" cellpadding="0" cellspacing="5">
    <tbody>
        <tr>
            <td width="35%" colspan="2" class="box-urutan-bg-004" style="color:white; text-align: center;">
                <b>GAYA PEKERJAAN</b>
            </td>
            <td width="10%" class="box-urutan-bg-004" style="color:white; text-align: center;">
                <b>SKOR</b>
            </td>
            <td width="50%" class="box-urutan-bg-004" style="color:white; text-align: center;">
                <b>KUALIFIKASI</b>
            </td>
        </tr>
        @foreach($gaya_pekerjaan as $r)
        <tr>
            <td  class="box-urutan-004" style="color:white; text-align: center;">
                {{$r->rangking}} 
            </td>
            <td  @if($r->rangking <=3) class="box-cell-info2" @else class="box-cell-info" @endif >
                {!! $r->cetak_komponen !!} 
            </td>
            <td  @if($r->rangking <=3) class="box-cell-info2" @else class="box-cell-info" @endif >
                <center>{{$r->skor}}</center>
            </td>
            <td @if($r->rangking <=3) class="box-cell-info2" @else class="box-cell-info" @endif style="padding-right:5px;" >
                @if($r->klasifikasi =="SR")
                    <div class="progress-bar2">
                        <span class="progress-bar-fill bg-fil-15" style="width: 10%;"> SANGAT RENDAH</span>
                    </div>
                @endif

                @if($r->klasifikasi =="RD")
                    <div class="progress-bar2">
                        <span class="progress-bar-fill bg-fil-20" style="width: 20%;"> RENDAH</span>
                    </div>
                @endif

                @if($r->klasifikasi =="AR")
                    <div class="progress-bar2">
                        <span class="progress-bar-fill bg-fil-30" style="width: 30%;"> AGAK RENDAH</span>
                    </div>
                @endif

                @if($r->klasifikasi =="SD")
                    <div class="progress-bar2">
                        <span class="progress-bar-fill bg-fil-55" style="width: 50%;"> SEDANG</span>
                    </div>
                @endif

                @if($r->klasifikasi =="AT")
                    <div class="progress-bar2">
                        <span class="progress-bar-fill bg-fill-70" style="width: 70%;"> AGAK TINGGI</span>
                    </div>
                @endif

                @if($r->klasifikasi =="TG")
                    <div class="progress-bar2">
                        <span class="progress-bar-fill bg-fill-85" style="width: 85%;"> TINGGI</span>
                    </div>
                @endif

                @if($r->klasifikasi =="ST")
                    <div class="progress-bar2">
                        <span class="progress-bar-fill bg-fil-100" style="width: 100%;"> SANGAT TINGGI</span>
                    </div>
                @endif
            </td>
        </tr>
        @endforeach
         
    </tbody>
</table>
