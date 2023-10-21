<?php 
loadHelper('skoring,function');
$def_kog = DB::table('ref_bidang_kognitif')->get();
$defenisi_kognitif = array();
foreach($def_kog as $d){
	$defenisi_kognitif[$d->field_skoring] = $d->deskripsi;
}
?>

<div  style="padding-left:5px;padding-right:5px;">
	<div class="box-title">
		<b>PROFIL KEMAMPUAN KOGNITIF</b>
	</div> 
	<br>
	<div class="box-subtitle">
		<b>Kecerdasan Umum (IQ) = {{$data_skoring->skor_iq}} ({{skoring_replace_(get_skor_predikat($data_skoring->tpa_iq,'skor_x','klasifikasi','ref_konversi_iq'))}})</b>
	</div>
</div>

<table width="100%" cellpadding="0" cellspacing="5">
	<tr>
		<td width="5%" class="box-cell-icon1" align="center" valign="align-middle">
			<img src="{{asset('images/icon/tpa_iu.png')}}" height="90px">
		</td>
		<td width="95%" class="box-cell-info">
				<b>INFORMASI UMUM</b> <br>
				{{ $defenisi_kognitif['tpa_iu']}}
				<?php 
				$klasifikasi = get_skor_predikat($data_skoring->tpa_iu,'skor','klasifikasi','ref_skala_kd_informasi_umum');
				?>
				@if($klasifikasi=='SANGAT_RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-20" style="width: 20%;"> SANGAT RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-40" style="width: 40%;"> RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='SEDANG')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-60" style="width: 60%;"> SEDANG</span>
				</div>
				@endif
				@if($klasifikasi=='TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-80" style="width: 80%;"> TINGGI</span>
				</div>
				@endif
				@if($klasifikasi=='SANGAT_TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-100" style="width: 100%;"> SANGAT TINGGI</span>
				</div>
				@endif
		</td>
	</tr>
	<tr>
		<td class="box-cell-icon1" align="center" valign="align-middle">
			<img src="{{asset('images/icon/tpa_pv.png')}}" height="90px">
		</td>
		<td class="box-cell-info">
				<b>PENALARAN VERBAL</b> <br>
				{{ $defenisi_kognitif['tpa_pv']}}
				<?php 
				$klasifikasi = get_skor_predikat($data_skoring->tpa_pv,'skor','klasifikasi','ref_skala_kd_penalaran_verbal');
				?>
				@if($klasifikasi=='SANGAT_RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-20" style="width: 20%;"> SANGAT RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-40" style="width: 40%;"> RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='SEDANG')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-60" style="width: 60%;"> SEDANG</span>
				</div>
				@endif
				@if($klasifikasi=='TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-80" style="width: 80%;"> TINGGI</span>
				</div>
				@endif
				@if($klasifikasi=='SANGAT_TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-100" style="width: 100%;"> SANGAT TINGGI</span>
				</div>
				@endif
		</td>
	</tr>

	<tr>
		<td class="box-cell-icon1" align="center" valign="align-middle">
			<img src="{{asset('images/icon/tpa_pk.png')}}" height="90px">
		</td>
		<td class="box-cell-info">
				<b>PENALARAN KUANTITATIF</b> <br>
				{{ $defenisi_kognitif['tpa_pk']}}
				<?php 
				$klasifikasi = get_skor_predikat($data_skoring->tpa_pk,'skor','klasifikasi','ref_skala_kd_penalaran_kuantitatif');
				?>
				@if($klasifikasi=='SANGAT_RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-20" style="width: 20%;"> SANGAT RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-40" style="width: 40%;"> RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='SEDANG')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-60" style="width: 60%;"> SEDANG</span>
				</div>
				@endif
				@if($klasifikasi=='TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-80" style="width: 80%;"> TINGGI</span>
				</div>
				@endif
				@if($klasifikasi=='SANGAT_TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-100" style="width: 100%;"> SANGAT TINGGI</span>
				</div>
				@endif
		</td>
	</tr>

	<tr>
		<td class="box-cell-icon1" align="center" valign="align-middle">
			<img src="{{asset('images/icon/tpa_pa.png')}}" height="90px">
		</td>
		<td class="box-cell-info">
				<b>PENALARAN ABSTRAK</b> <br>
				{{ $defenisi_kognitif['tpa_pa']}}
				<?php 
				$klasifikasi = get_skor_predikat($data_skoring->tpa_pa,'skor','klasifikasi','ref_skala_kd_penalaran_abstrak');
				?>
				@if($klasifikasi=='SANGAT_RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-20" style="width: 20%;"> SANGAT RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-40" style="width: 40%;"> RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='SEDANG')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-60" style="width: 60%;"> SEDANG</span>
				</div>
				@endif
				@if($klasifikasi=='TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-80" style="width: 80%;"> TINGGI</span>
				</div>
				@endif
				@if($klasifikasi=='SANGAT_TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-100" style="width: 100%;"> SANGAT TINGGI</span>
				</div>
				@endif
		</td>
	</tr>

	<tr>
		<td class="box-cell-icon1" align="center" valign="align-middle">
			<img src="{{asset('images/icon/tpa_ps.png')}}" height="90px">
		</td>
		<td class="box-cell-info">
				<b>PENALARAN SPASIAL</b> <br>
				{{ $defenisi_kognitif['tpa_ps']}}
				<?php 
				$klasifikasi = get_skor_predikat($data_skoring->tpa_ps,'skor','klasifikasi','ref_skala_kd_penalaran_spasial');
				?>
				@if($klasifikasi=='SANGAT_RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-20" style="width: 20%;"> SANGAT RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-40" style="width: 40%;"> RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='SEDANG')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-60" style="width: 60%;"> SEDANG</span>
				</div>
				@endif
				@if($klasifikasi=='TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-80" style="width: 80%;"> TINGGI</span>
				</div>
				@endif
				@if($klasifikasi=='SANGAT_TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-100" style="width: 100%;"> SANGAT TINGGI</span>
				</div>
				@endif
		</td>
	</tr>

	<tr>
		<td class="box-cell-icon1" align="center" valign="align-middle">
			<img src="{{asset('images/icon/tpa_pm.png')}}" height="90px">
		</td>
		<td class="box-cell-info">
				<b>PENGERTIAN MEKANIKA</b> <br>
				{{ $defenisi_kognitif['tpa_pm']}}
				<?php 
				$klasifikasi = get_skor_predikat($data_skoring->tpa_pm,'skor','klasifikasi','ref_skala_kd_penalaran_mekanika');
				?>
				@if($klasifikasi=='SANGAT_RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-20" style="width: 20%;"> SANGAT RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-40" style="width: 40%;"> RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='SEDANG')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-60" style="width: 60%;"> SEDANG</span>
				</div>
				@endif
				@if($klasifikasi=='TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-80" style="width: 80%;"> TINGGI</span>
				</div>
				@endif
				@if($klasifikasi=='SANGAT_TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-100" style="width: 100%;"> SANGAT TINGGI</span>
				</div>
				@endif
		</td>
	</tr>

	<tr>
		<td class="box-cell-icon1" align="center" valign="align-middle">
			<img src="{{asset('images/icon/tpa_kt.png')}}" height="90px">
		</td>
		<td class="box-cell-info">
				<b>KETELITIAN</b> <br>
				{{ $defenisi_kognitif['tpa_kt']}}
				<?php 
				$klasifikasi = get_skor_predikat($data_skoring->tpa_kt,'skor','klasifikasi','ref_skala_kd_cepat_teliti');
				?>
				@if($klasifikasi=='SANGAT_RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-20" style="width: 20%;"> SANGAT RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='RENDAH')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-40" style="width: 40%;"> RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='SEDANG')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-60" style="width: 60%;"> SEDANG</span>
				</div>
				@endif
				@if($klasifikasi=='TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-80" style="width: 80%;"> TINGGI</span>
				</div>
				@endif
				@if($klasifikasi=='SANGAT_TINGGI')
				<div class="progress-bar">
					<span class="progress-bar-fill bg-fil-100" style="width: 100%;"> SANGAT TINGGI</span>
				</div>
				@endif
 
		</td>
	</tr>

</table>
