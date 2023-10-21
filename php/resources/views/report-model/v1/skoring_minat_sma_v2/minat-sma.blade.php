<?php 
loadHelper('skoring,function');
$list_klasifikasi= get_list_enum_values('ref_klasifikasi_minat_sma', 'klasifikasi');
//var_dump($list_klasifikasi);
?>
<div  style="padding-left:5px;padding-right:5px;">
	<div class="box-title-008">
		<b>PEMINATAN SMA</b>
	</div> 
	<br>
	<div class="box-subtitle-008">
		<b>Klasifikasi Tingkat Peminatan SMA Berdasarkan 3 Kelompok Minat</b>
	</div>
</div>
<?php
//ref_klasifikasi_minat_sma
$list = DB::select("SELECT * FROM ref_pilihan_minat_sma order by kd_pilihan");
//var_dump($list);
?>
<table width="100%" cellpadding="0" cellspacing="5">
	@foreach($list as $r)
	<tr>
		<td width="5%" class="box-cell-icon2  box-icon-008" align="center" valign="align-middle">
			<img src="{{asset('images/icon/'.$r->gambar)}}" height="90px">
		</td>
		<td width="95%" class="box-cell-info">
				<b>{{skoring_replace_($r->nama_pilihan)}}</b> <br>
				{{$r->keterangan}}
				<?php 
				$field_skoring = $r->field_skoring;
				$klasifikasi = get_skor_predikat($data_skoring->$field_skoring,'skor','klasifikasi','ref_klasifikasi_minat_sma');
				?>
				@if($klasifikasi=='SANGAT_RENDAH')
				<div class="progress-bar-008">
					<span class="progress-bar-fill bg-fil-008-20" style="width: 20%;"> SANGAT RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='RENDAH')
				<div class="progress-bar-008">
					<span class="progress-bar-fill bg-fil-008-40" style="width: 40%;"> RENDAH</span>
				</div>
				@endif
				@if($klasifikasi=='SEDANG')
				<div class="progress-bar-008">
					<span class="progress-bar-fill bg-fil-008-60" style="width: 60%;"> SEDANG</span>
				</div>
				@endif
				@if($klasifikasi=='TINGGI')
				<div class="progress-bar-008">
					<span class="progress-bar-fill bg-fil-008-80" style="width: 80%;"> TINGGI</span>
				</div>
				@endif
				@if($klasifikasi=='SANGAT_TINGGI')
				<div class="progress-bar-008">
					<span class="progress-bar-fill bg-fil-008-100" style="width: 100%;"> SANGAT TINGGI</span>
				</div>
				@endif
 
		</td>
	</tr>
	@endforeach
</table>
 