<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge"> 
  </head>

  <style>
    
    .facet-list {
    list-style-type: none;
    margin: 0;
    padding: 0;
    margin-top: 20px;
    margin-right: 10px;  
    width: 100%;
    min-height: 1.5em;
    font-size: 0.85em;
    }
    .facet-list li {
    margin: 5px;
    padding: 5px 15px 5px 15px;
    font-size: 1.2em;
    width: 100%%;
    }
    .facet-list li.placeholder {
    height: 1.2em
    }
    .facet {
    border: 1px solid rgb(146, 143, 152);
    background-color: #f5f4f5;
    cursor: move;
    }
    .facet.ui-sortable-helper {
    opacity: 0.5;
    }
    .placeholder { 
    background-color: #e1d5ea;
    }
    </style>
  <body>
    @if(count($list)==0)
    <br>
    <p class="facet" style="padding: 10px;">Belum ada komponen laporan</p>
    @endif
    <ul id="list-report" class="facet-list">
        @foreach($list as $r)
        <li class="facet" style=" @if($r->tabel_referensi != "-") background:#fff !important;  @endif" id="{{$r->uuid}}">
            <div class="row">
                <div class="col-md-9">
                    @if($r->tabel_referensi != "-")
                        <div>{{$r->nama_report}}</div>
                    @else 
                        <div><small>{{$r->nama_report}}</small></div>
                    @endif
                </div>
                <div class="col-md-3 d-flex flex-row-reverse ">
                      <button type="button" data-uuid="{{$r->uuid}}" 
                        class="btn btn-konfirm-delete  btn-outline-danger btn-sm"><i class="la la-times"></i></button> 
                    @if($r->tabel_referensi != "-")
                        &nbsp;&nbsp;
                        <button type="button" data-uuid="{{$r->uuid}}" 
                          class="btn btn-preview-komponen  btn-outline-secondary btn-sm"><i class="la la-search-plus"></i></button> 
                    @endif
                </div>
            </div> 
        </li>
        @endforeach
    </ul>
  </body>
</html>