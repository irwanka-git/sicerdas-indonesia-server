<?php 
loadHelper('skoring,function');
?>
<div  style="padding-left:5px;padding-right:5px;">
	<div class="box-title-009">
		<b>SIKAP TERHADAP PELAJARAN</b>
	</div> 
	<br>
	<div class="box-subtitle-009">
		<b>Klasifikasi Sikap Terhadap Pelajaran Berdasarkan 4 Kelompok Pelajaran</b>
	</div>
</div>
<table width="100%" cellpadding="0" cellspacing="5">
	<?php
	$list_klasifikasi= get_list_enum_values('ref_skala_sikap_pelajaran', 'klasifikasi');
	?>
	<tbody>
		<?php 
		$list = DB::table('soal_sikap_pelajaran_kuliah')->orderby('urutan','asc')->get();
		$item_kelompok['Ilmu Dasar'] = 4;
		$item_kelompok['Ilmu Alam'] = 4;
		$item_kelompok['Ilmu Sosial'] = 4;
		$item_kelompok['Ilmu Praktis'] = 4;
		$kelompok = '';
		?>
		@foreach($list as $r)
		<tr>
			<?php 
			if($r->kelompok!=$kelompok){
				$kelompok = $r->kelompok;
				$gambar = str_replace(" ","", strtolower($kelompok))
			?>
				<td rowspan="{{$item_kelompok[$kelompok]}}" class="box-urutan-009" align="center" width="20%">
					<img src="{{asset('images/icon/'.$gambar.'.png')}}" height="80px"><br>
					{{skoring_replace_($kelompok)}}
				</td>
			<?php 
			}
			?>
			<td width="30%" class="box-cell-info-009">{{$r->pelajaran}}</td>
			<?php 
			$field_skoring = $r->field_skoring;
			$klasifikasi = get_skor_predikat($data_skoring->$field_skoring,'skor','klasifikasi','ref_skala_sikap_pelajaran');
			?>
			<td>
			@if($klasifikasi=='Netral')
			<div class="progress-bar-009">
				<span class="progress-bar-fill-009 bg-fil-009-netral" style="width: 20%;"> Netral</span>
			</div>
			@endif
			@if($klasifikasi=='Sangat Negatif')
			<div class="progress-bar-009">
				<span class="progress-bar-fill-009 bg-fil-009-sangat-negatif" style="width: 50%;"> Sangat Negatif</span>
			</div>
			@endif
			@if($klasifikasi=='Negatif')
			<div class="progress-bar-009">
				<span class="progress-bar-fill-009 bg-fil-009-negatif" style="width: 30%;"> Negatif</span>
			</div>
			@endif
			@if($klasifikasi=='Positif')
			<div class="progress-bar-009">
				<span class="progress-bar-fill-009 bg-fil-009-positif" style="width: 30%;"> Positif</span>
			</div>
			@endif
			@if($klasifikasi=='Sangat Positif')
			<div class="progress-bar-009">
				<span class="progress-bar-fill-009 bg-fil-009-sangat-positif" style="width: 50%;"> Sangat Positif</span>
			</div>
			@endif
			</td>
		</tr>
		@endforeach
	</tbody>
</table>