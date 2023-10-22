<?php

 
namespace App\Http\Middleware;

use Closure;
use DB;
use URL;

//use Illuminate\Http\Request;
class FilterHttps
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
            URL::forceScheme("https");
        }
        if ($request->root()=='http://www.sicerdas.web.id'){
            URL::forceScheme("https");
        }
        return $next($request);
    }
 

}
