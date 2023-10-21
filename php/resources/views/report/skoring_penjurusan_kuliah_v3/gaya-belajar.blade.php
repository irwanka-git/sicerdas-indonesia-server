<?php 
loadHelper('skoring,function'); 
//var_dump($list_klasifikasi);
?>
<table class="table table-striped table-sm table-bordered table-x">
	<thead>
		<tr class="table-primary">
			<th colspan="5" class="text-center sm">SKALA GAYA BELAJAR</th>
		</tr>
		<tr class="table-primary">
			<th width="45%" class="text-center sm">GAYA BELAJAR</th>
			<th class="text-center sm">SKOR</th>
			<th class="text-center sm">SEDIKIT</th>
			<th class="text-center sm">SEDANG</th>
			<th class="text-center sm">DOMINAN</th>
		</tr>
	</thead>
	<tbody>
		<?php
		//ref_klasifikasi_minat_sma
		$list = DB::select("SELECT * FROM ref_gaya_belajar order by kode");
		//var_dump($list);
		?>
		@foreach($list as $r)
		<tr>
			<td class="jf">
				{{strtoupper($r->nama)}}<br>
				<small>{{$r->deskripsi}}</small>
			</td>
			<?php 
			$field_name = "gaya_".strtolower($r->nama);
			$klasifikasi = get_skor_predikat($data_skoring->$field_name,'skor','klasifikasi','ref_skor_gaya_belajar');
			?>
			<td align="center">
				{{$data_skoring->$field_name}}
			</td>
			<td align="center">
				@if($klasifikasi=="SEDIKIT") <i class="fas fa-check"></i> @endif
			</td>
			<td align="center">
				@if($klasifikasi=="SEDANG") <i class="fas fa-check"></i> @endif
			</td>
			<td align="center">
				@if($klasifikasi=="DOMINAN") <i class="fas fa-check"></i> @endif
			</td>
		</tr>
		@endforeach
	</tbody>
</table>