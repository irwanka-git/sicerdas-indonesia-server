<?php

namespace App\Http\Middleware;

use Closure;
use DB;

//use Illuminate\Http\Request;
class HandleHttp
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle($request, Closure $next)
    {   
         
        if ($request->root()=='http://sicerdas.web.id'){
            URL::forceSchema("https");
        }
        return $next($request);
    }
 

}