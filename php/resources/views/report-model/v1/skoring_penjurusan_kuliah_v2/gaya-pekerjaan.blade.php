<?php 
loadHelper('skoring,function');
$id_user = $data_skoring->id_user;
$id_quiz = $data_skoring->id_quiz;

$list_klasifikasi= DB::table('ref_klasifikasi_gaya_kerja')->select('akronim')->get();
$gaya_pekerjaan = DB::select("SELECT A
    .ID,
    A.rangking,
    C.nama_komponen,
    C.deskripsi,
    C.kode,
    C.cetak_komponen,
    A.skor,
    A.klasifikasi, 
    C.gambar
FROM
    ref_skoring_gaya_pekerjaan AS A,
    ref_komponen_gaya_pekerjaan AS C 
WHERE
    C.kode = A.kode 
    AND A.id_quiz = $id_quiz 
    AND A.id_user = $id_user 
ORDER BY
    A.rangking limit 3");

?>
<div  style="padding-left:5px;padding-right:5px;">
	<div class="box-title-300">
		<b>GAYA PEKERJAAN</b>
	</div> 
	<br>
	<div class="box-subtitle-300">
		<b>Keterangan: Urutan nomor 1 (satu) sampai 3 (tiga) yang paling dominan dari 12 (dua belas) gaya pekerjaan
        </b>
	</div>
</div>
<table width="100%" cellpadding="0" cellspacing="5">
    <tbody>

        @foreach($gaya_pekerjaan as $r)
         <tr>
            <td width="5%" class="box-urutan-300">
                {{$r->rangking}} 
            </td>
            <td width="5%" class="box-cell-icon2-300  box-icon-180" align="center" valign="align-middle">
                <img src="{{asset('images/icon/'.$r->gambar)}}" height="90px">
            </td>
             <td class="sm box-cell-info">  
                <b>{!! strtoupper($r->cetak_komponen) !!}</b><br>
                <small>{{$r->deskripsi}}</small>
                @if($r->klasifikasi =="SR")
				<div class="progress-bar-300">
					<span class="progress-bar-fill bg-fil-300-1" style="width: 10%;"> SANGAT RENDAH ({{$r->skor}})</span>
				</div>
                @endif

                @if($r->klasifikasi =="RD")
				<div class="progress-bar-300">
					<span class="progress-bar-fill bg-fil-300-2" style="width: 20%;"> RENDAH ({{$r->skor}})</span>
				</div>
                @endif

                @if($r->klasifikasi =="AR")
				<div class="progress-bar-300">
					<span class="progress-bar-fill bg-fil-300-3" style="width: 35%;"> AGAK RENDAH ({{$r->skor}})</span>
				</div>
                @endif

                @if($r->klasifikasi =="SD")
				<div class="progress-bar-300">
					<span class="progress-bar-fill bg-fil-300-4" style="width: 50%;"> SEDANG ({{$r->skor}})</span>
				</div>
                @endif

                @if($r->klasifikasi =="AT")
				<div class="progress-bar-300">
					<span class="progress-bar-fill bg-fil-300-5" style="width: 70%;"> AGAK TINGGI ({{$r->skor}})</span>
				</div>
                @endif

                @if($r->klasifikasi =="TG")
				<div class="progress-bar-300">
					<span class="progress-bar-fill bg-fil-300-6" style="width: 80%;"> TINGGI ({{$r->skor}})</span>
				</div>
                @endif

                @if($r->klasifikasi =="ST")
				<div class="progress-bar-300">
					<span class="progress-bar-fill bg-fil-300-7" style="width: 100%;"> SANGAT TINGGI ({{$r->skor}})</span>
				</div>
                @endif
				 
            </td>
         </tr>
        @endforeach
    </tbody>
</table>
