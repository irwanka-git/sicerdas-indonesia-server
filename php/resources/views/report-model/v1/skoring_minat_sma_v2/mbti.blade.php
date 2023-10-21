<?php
loadHelper('skoring,function');
?>
<table width="100%" cellpadding="0" cellspacing="5">
	<thead>
		<tr >
			<th width="100%" class="box-header-011">TIPOLOGI KEPRIBADIAN</th>
		</tr>
	</thead>
	<tbody>
		<?php 
			$list = DB::table('interprestasi_tipologi_jung')->where('kode',$data_skoring->tipojung_kode)->first();
		?>
            @if($list)
			<tr>

				<td class="box-cell-info-011">
					<strong>{{ $list->urutan }} - {{ $list->kode }} - {{ $list->nama }} -  <i>{{ $list->keterangan }}</i></strong>
	                <br>
	                <p style="text-align: justify">{{ $list->deskripsi }}</p>
	        
	                <b style="text-align: justify">Informasi lebih lanjut bisa dilihat di internet<i> <a href="https://www.google.com/search?q={{ $list->keterangan }}" target="_blank">"{{ $list->keterangan }}"</a></i></b>
				</td>
			</tr>
			@else
			<tr>
				<td class="box-cell-info-011">
					{{$data_skoring->tipojung_kode }} Tidak Ditemukan
				</td>
			</tr>
			@endif
	</tbody>
</table>