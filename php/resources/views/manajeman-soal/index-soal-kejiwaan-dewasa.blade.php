<?php
$main_path = Request::segment(1);
loadHelper('akses'); 
?>
@extends('layout')
@section("css")
<link rel="stylesheet" href="//cdn.quilljs.com/1.3.6/quill.snow.css" />
@endsection
@section("pagetitle")
	 
@endsection

@section('content')

		<div class="row">
			<div class="col-12">
				<div class="card">
					<div class="card-header">
						<h5 class="card-title">{{$pagetitle}}</h5>
						<h6 class="card-subtitle text-muted">
                            {!! $smalltitle !!}
                         </h6>
					</div>
					<div class="card-body">
						<table id="datatable" class="table table-striped table-hover table-sm" style="width:100%">
							<thead>
								<tr>
									<th width="5%">#</th>									
									<th width="15%">Model</th>
									<th width="5%">Urutan</th>
									<th>Unsur</th> 
								</tr>
							</thead>
							<tbody>
								 
							</tbody>
						</table>


					</div>
				</div>
			</div>
		</div>
 
@endsection

@section("modal")
 
  
@endsection

@section("js")
<script src="//cdn.quilljs.com/1.3.6/quill.min.js"></script>
<script type="text/javascript">
	$(function(){

			var toolbarOptions = [['bold', 'italic'], ['link', 'image']];
			var $tabel1 = $('#datatable').DataTable({
			    processing: true,
			    responsive: true,
			    fixedHeader: true,
			    serverSide: true,
			    ajax: "{{url('soal-kejiwaan-dewasa/dt')}}",
			    "iDisplayLength": 25,
			    columns: [
			    	 {data:'DT_Row_Index' , orderable:false, searchable: false,sClass:""},
			         {data:'model' , name:"model" , orderable:false, searchable: false,sClass:""},
			         {data:'urutan' , name:"urutan" , orderable:false, searchable: false,sClass:"text-center"},
			         {data:'unsur' , name:"unsur" , orderable:false, searchable: false,sClass:""},
			        ],
			        "fnRowCallback": function( nRow, aData, iDisplayIndex, iDisplayIndexFull ) {
			        $(nRow).addClass( aData["rowClass"] );
			        return nRow;
			    },
			     
			});
 
	})
</script>
@endsection