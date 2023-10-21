<?php 
loadHelper('skoring,function'); 
?>
<div  style="padding-left:5px;padding-right:5px;">
	<div class="box-title-002">
		<b>PEMINATAN SMK</b>
	</div> 
	<br>
	<div class="box-subtitle-002">
        <small>Keterangan: Urutan nomor 1 (satu) adalah jurusan kuliah yang paling dominan, urutan momor 2 (dua) adalah jurusan kuliah kedua yang dominan. Begitupun yang seterusnya.</small>
	</div>
</div>
<table width="100%" cellpadding="0" cellspacing="5">
	<tbody>
		<?php
		//ref_klasifikasi_minat_sma
		$arr_minat = [];
		$list = DB::select("SELECT * FROM soal_peminatan_smk");
		$jumlah_pilihan = DB::table('quiz_sesi_mapping_smk')->where('id_quiz', $data_sesi->id_quiz)->count();
		foreach ($list as $l){
			$arr_minat[$l->nomor]['keterangan'] = $l->keterangan;
			$arr_minat[$l->nomor]['deskripsi'] = $l->deskripsi;
			$arr_minat[$l->nomor]['gambar'] = $l->gambar;
		}
		?>
		<?php 
		for ($n=1;$n<=$jumlah_pilihan;$n++){
			$val_minat = 'minat_'.$n;
		?>
			<tr>
				<td width="5%" class="box-urutan-002">
					{{$n}}
				</td>
				<td width="10%" class="box-cell-icon2 box-icon-002" align="center" valign="align-middle">
					<img src="{{asset('gambar/'.$arr_minat[$data_skoring->$val_minat]['gambar'])}}" height="90px">
				</td>
				<td width="85%" class="box-cell-info">
					<b>{{strtoupper($arr_minat[$data_skoring->$val_minat]['keterangan'])}}</b><br>
					<small>{!! strip_tags($arr_minat[$data_skoring->$val_minat]['deskripsi'], '<b><strong>') !!}</small>
				</td>
			</tr>
		<?php 
		}
		?>
		
	</tbody>
</table>
