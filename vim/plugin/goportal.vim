if !exists('g:goportal_default_fifoname')
    let g:goportal_default_fifoname = ""
endif

function! GoPortal(...)
    let cmd = join(a:000," ")
    call system("goportal ".cmd)
endfunction

function! GoPortalDefault(...)
    if a:1 == ""
        echom g:goportal_default_fifoname
    else
        let g:goportal_default_fifoname = a:1
    endif
endfunction

command! -nargs=* -complete=file Gp   call GoPortal(g:goportal_default_fifoname,<q-args>)
command! -nargs=+ -complete=file Gpi  call GoPortal("-i",g:goportal_default_fifoname,<q-args>)
command! -nargs=+ -complete=file Gpo  call GoPortal("-o",g:goportal_default_fifoname,<q-args>)
command! -nargs=* Gpd  call GoPortalDefault(<q-args>)
command! -nargs=0 Gpdd  call GoPortalDefault(" ")
