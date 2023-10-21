<?php 
loadHelper('skoring,function');
$list_klasifikasi= get_list_enum_values('ref_klasifikasi_minat_sma', 'klasifikasi');
//var_dump($list_klasifikasi);
?>
<div  style="padding-left:5px;padding-right:5px;">
	<div class="box-title-180">
		<b>GAYA BELAJAR</b>
	</div> 
	<br>
	<div class="box-subtitle-180">
		<b>Keterangan:  Skala Gaya Belajar Berdasarkan 3 Kelompok</b>
	</div>
</div>
<?php 
$list = DB::select("SELECT * FROM ref_gaya_belajar order by kode"); 
?>
<table width="100%" cellpadding="0" cellspacing="5">
	@foreach($list as $r)
	<tr>
		<td width="5%" class="box-cell-icon2-180  box-icon-180" align="center" valign="align-middle">
			<img src="{{asset('images/icon/'.$r->gambar)}}" height="90px">
		</td>
		<td width="95%" class="box-cell-info">
				<b>{{strtoupper($r->nama)}}</b> <br>
				{{$r->deskripsi}}
				<?php 
				$field_skoring = "gaya_".strtolower($r->nama);
				$klasifikasi = get_skor_predikat($data_skoring->$field_skoring,'skor','klasifikasi','ref_skor_gaya_belajar');
				?>
				@if($klasifikasi=='SEDIKIT')
				<div class="progress-bar-180">
					<span class="progress-bar-fill bg-fil-180-20" style="width: 20%;"> SEDIKIT ({{$data_skoring->$field_skoring}})</span>
				</div>
				@endif
                @if($klasifikasi=='SEDANG')
				<div class="progress-bar-180">
					<span class="progress-bar-fill bg-fil-180-60" style="width: 50%;"> SEDANG ({{$data_skoring->$field_skoring}})</span>
				</div>
				@endif
				@if($klasifikasi=='DOMINAN')
				<div class="progress-bar-180">
					<span class="progress-bar-fill bg-fil-180-100" style="width: 100%;"> DOMINAN ({{$data_skoring->$field_skoring}})</span>
				</div>
				@endif
 
		</td>
	</tr>
	@endforeach
</table>
 
 