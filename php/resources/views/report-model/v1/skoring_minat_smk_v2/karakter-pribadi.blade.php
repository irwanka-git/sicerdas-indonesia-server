<?php 
loadHelper('skoring,function');
$list_klasifikasi= get_list_enum_values('ref_karakter_pribadi', 'klasifikasi'); 
$singkat_nama = array(
	"Sangat Rendah"=>"SR",
	"Rendah"=>"R",
	"Sedang"=>"S",
	"Tinggi"=>"T",
	"Sangat Tinggi"=>"ST",
	);
//dd($list_klasifikasi);
?>
<div  style="padding-left:5px;padding-right:5px;">
	<div class="box-title-007">
		<b>KARAKTERISTIK PRIBADI</b>
	</div> 
	<br>
	<div class="box-subtitle-007">
		<b>Terdapat 12 Komponen Karakteristik Pribadi</b>
	</div>
</div>


<table width="100%" cellpadding="0" cellspacing="5">
	<?php
	$no=1;
	$komponen = DB::table('ref_komponen_karakteristik_pribadi')->orderby('id_komponen','asc')->get();
	?>
	@foreach($komponen as $k)
	<tr>
		<td width="5%" class="box-cell-icon1 box-icon-007" align="center" valign="align-middle">
			<img src="{{asset('images/icon/'.$k->icon)}}" height="90px">
		</td>
		<td width="95%" class="box-cell-info">
			<b>{{strtoupper($k->nama_komponen)}}</b><br>
			<small>{{$k->keterangan}}</small>
			<?php 
			$field_skoring = $k->field_skoring;
			$klasifikasi = get_skor_predikat($data_skoring->$field_skoring,'skor','klasifikasi','ref_karakter_pribadi');
			?>
			@if($klasifikasi=='Sangat Rendah')
			<div class="progress-bar-blue">
				<span class="progress-bar-fill bg-fil2-20" style="width: 20%;"> SANGAT RENDAH</span>
			</div>
			@endif
			@if($klasifikasi=='Rendah')
			<div class="progress-bar-blue">
				<span class="progress-bar-fill bg-fil2-40" style="width: 40%;"> RENDAH</span>
			</div>
			@endif
			@if($klasifikasi=='Sedang')
			<div class="progress-bar-blue">
				<span class="progress-bar-fill bg-fil2-60" style="width: 60%;"> SEDANG</span>
			</div>
			@endif
			@if($klasifikasi=='Tinggi')
			<div class="progress-bar-blue">
				<span class="progress-bar-fill bg-fil2-80" style="width: 80%;"> TINGGI</span>
			</div>
			@endif
			@if($klasifikasi=='Sangat Tinggi')
			<div class="progress-bar-blue">
				<span class="progress-bar-fill bg-fil2-100" style="width: 100%;"> SANGAT TINGGI</span>
			</div>
			@endif
			
		</td>
	</tr>
	@endforeach
</table>
 