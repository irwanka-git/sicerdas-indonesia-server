@extends('report.layout')
@section('content')
<?php loadHelper('skoring,function'); 
	$cek_nomor_seri = env('APP_URL').'/vcf/'.$nomor_seri;
?>
<link class="stylesheet" href="{{url('css/report/model.v1.css')}}" rel="stylesheet">
<?php 
$biro = DB::table('users')->where('id', $data_sesi->id_user_biro)->first();
$lokasi = DB::table('lokasi')->where('id_lokasi', $quiz->id_lokasi)->first();
?>
	<img class="kop" src="{{'/kop/'.$biro->kop_biro}}" height="80px">

	<center style="margin-top: 30px;">
		<b style="font-size: 18px;">P S I K O G R A M</b><br>
		<small>{{$data_sesi->nama_sesi}}</small>
	</center>

	<div style="margin-bottom: 20px; margin-top: 30px;">
		<table width="100%">
			<tr>
				<td class="box-cell-icon1 bg-001" width="5%" align="center" valign="align-middle">
					<img src="{{asset('images/icon/user.png')}}" height="90px">
				</td>
				<td class="box-cell-info" width="95%" style="padding-top: 10px;padding-bottom: 10px;">
						<span style="float:right;">
							{!! DNS2D::getBarcodeSVG($cek_nomor_seri, 'QRCODE',3,3); !!}
						</span>

						<b>{{$user->nama_pengguna}}</b><br>
						{{($user->jenis_kelamin=='L' || $user->jenis_kelamin=='M') ? 'Laki-Laki' : 'Perempuan'}}
						<br>
						{{$user->unit_organisasi}}  /  {{$user->organisasi}}<br>
						Tanggal Tes :  {{tgl_indo_lengkap(substr($data_sesi->tanggal,0,10))}}<br>
						Lokasi Tes :  {{$lokasi->nama_lokasi}}<br>
				</td>
		</tr>
		</table>
	</div>
	@include("report-model.v1.".$data_sesi->skoring_tabel.".kognitif")
	<p class="new-page"></p>
	@include("report-model.v1.".$data_sesi->skoring_tabel.".minat-sma")
	<br> 
	@include("report-model.v1.".$data_sesi->skoring_tabel.".sikap-pelajaran") 
	<p class="new-page"></p>
	@include("report-model.v1.".$data_sesi->skoring_tabel.".minat-sains")
	<br>
	@include("report-model.v1.".$data_sesi->skoring_tabel.".minat-sosial")
	<p class="new-page"></p>
	@include("report-model.v1.".$data_sesi->skoring_tabel.".rekomendasi-jurusan-kuliah")
	<br>
	@include("report-model.v1.".$data_sesi->skoring_tabel.".mbti")
	<br>
	@include("report-model.v1.".$data_sesi->skoring_tabel.".karakter-pribadi")
	
	<p class="new-page"></p> 
	@include("report-model.v1.".$data_sesi->skoring_tabel.".saran")
	@if($data_sesi->ttd_asesor)
		<?php 
		$cek_ttd = DB::table('gambar')->where('filename', $data_sesi->ttd_asesor)->first();
		$ttd_gambar =  '<img width="125px" src="data:image/'.$cek_ttd->type.';base64,'.$cek_ttd->image_base64. '">';
		?>
			<center>
				{{$data_sesi->kota}}, {{tgl_indo(substr($data_sesi->tanggal,0,10))}}<br>
				Asesor,
					<br>{!! $ttd_gambar !!}<br>
				 {{$data_sesi->nama_asesor}}<br>
				 <small>SIPP: {{$data_sesi->nomor_sipp}}</small>
			</center>
	  @else
		  <center>
				{{$data_sesi->kota}}, {{tgl_indo(substr($data_sesi->tanggal,0,10))}}<br>
				Asesor,<br><br><br><br>
				 {{$data_sesi->nama_asesor}}<br>
				 <small>SIPP: {{$data_sesi->nomor_sipp}}</small>
		 </center>
	  @endif
@endsection