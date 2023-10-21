<?php
$quiz_template_saran = DB::table('quiz_template_saran')->where('skoring_tabel', $data_sesi->skoring_tabel)->first(); 
?>
<table class="table table-striped table-sm table-bordered  table-x">
	<thead>
		<tr class="table-primary">
			<th class="box-header sm">SARAN</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td  class="jf box-cell-info"  style="padding-left:20px; padding-right: 20px; line-height: 1.15em !important; font-size: 0.9em !important">{!! $quiz_template_saran->isi !!}</td>
		</tr>
	</tbody>
</table>
  