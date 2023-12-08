(defn is-prime? [num]
  (if (< num 2)
    false
    (loop [i 2]
      (if (<= (* i i) num)
        (if (zero? (mod num i))
          false
          (recur (inc i)))
        true))))

(defn check-prime [num]
  (cond
    (= num  1 ) (is-prime?  1 )
    (= num  2 ) (is-prime?  2 )
    (= num  3 ) (is-prime?  3 )
    (= num  4 ) (is-prime?  4 )
    (= num  5 ) (is-prime?  5 )
    (= num  6 ) (is-prime?  6 )
    (= num  7 ) (is-prime?  7 )
    (= num  8 ) (is-prime?  8 )
    (= num  9 ) (is-prime?  9 )
    (= num  10 ) (is-prime?  10 )
    (= num  11 ) (is-prime?  11 )
    (= num  12 ) (is-prime?  12 )
    (= num  13 ) (is-prime?  13 )
    (= num  14 ) (is-prime?  14 )
    (= num  15 ) (is-prime?  15 )
    (= num  16 ) (is-prime?  16 )
    (= num  17 ) (is-prime?  17 )
    (= num  18 ) (is-prime?  18 )
    (= num  19 ) (is-prime?  19 )
    (= num  20 ) (is-prime?  20 )
    (= num  21 ) (is-prime?  21 )
    (= num  22 ) (is-prime?  22 )
    (= num  23 ) (is-prime?  23 )
    (= num  24 ) (is-prime?  24 )
    (= num  25 ) (is-prime?  25 )
    (= num  26 ) (is-prime?  26 )
    (= num  27 ) (is-prime?  27 )
    (= num  28 ) (is-prime?  28 )
    (= num  29 ) (is-prime?  29 )
    (= num  30 ) (is-prime?  30 )
    (= num  31 ) (is-prime?  31 )
    (= num  32 ) (is-prime?  32 )
    (= num  33 ) (is-prime?  33 )
    (= num  34 ) (is-prime?  34 )
    (= num  35 ) (is-prime?  35 )
    (= num  36 ) (is-prime?  36 )
    (= num  37 ) (is-prime?  37 )
    (= num  38 ) (is-prime?  38 )
    (= num  39 ) (is-prime?  39 )
    (= num  40 ) (is-prime?  40 )
    (= num  41 ) (is-prime?  41 )
    (= num  42 ) (is-prime?  42 )
    (= num  43 ) (is-prime?  43 )
    (= num  44 ) (is-prime?  44 )
    (= num  45 ) (is-prime?  45 )
    (= num  46 ) (is-prime?  46 )
    (= num  47 ) (is-prime?  47 )
    (= num  48 ) (is-prime?  48 )
    (= num  49 ) (is-prime?  49 )
    (= num  50 ) (is-prime?  50 )
    (= num  51 ) (is-prime?  51 )
    (= num  52 ) (is-prime?  52 )
    (= num  53 ) (is-prime?  53 )
    (= num  54 ) (is-prime?  54 )
    (= num  55 ) (is-prime?  55 )
    (= num  56 ) (is-prime?  56 )
    (= num  57 ) (is-prime?  57 )
    (= num  58 ) (is-prime?  58 )
    (= num  59 ) (is-prime?  59 )
    (= num  60 ) (is-prime?  60 )
    (= num  61 ) (is-prime?  61 )
    (= num  62 ) (is-prime?  62 )
    (= num  63 ) (is-prime?  63 )
    (= num  64 ) (is-prime?  64 )
    (= num  65 ) (is-prime?  65 )
    (= num  66 ) (is-prime?  66 )
    (= num  67 ) (is-prime?  67 )
    (= num  68 ) (is-prime?  68 )
    (= num  69 ) (is-prime?  69 )
    (= num  70 ) (is-prime?  70 )
    (= num  71 ) (is-prime?  71 )
    (= num  72 ) (is-prime?  72 )
    (= num  73 ) (is-prime?  73 )
    (= num  74 ) (is-prime?  74 )
    (= num  75 ) (is-prime?  75 )
    (= num  76 ) (is-prime?  76 )
    (= num  77 ) (is-prime?  77 )
    (= num  78 ) (is-prime?  78 )
    (= num  79 ) (is-prime?  79 )
    (= num  80 ) (is-prime?  80 )
    (= num  81 ) (is-prime?  81 )
    (= num  82 ) (is-prime?  82 )
    (= num  83 ) (is-prime?  83 )
    (= num  84 ) (is-prime?  84 )
    (= num  85 ) (is-prime?  85 )
    (= num  86 ) (is-prime?  86 )
    (= num  87 ) (is-prime?  87 )
    (= num  88 ) (is-prime?  88 )
    (= num  89 ) (is-prime?  89 )
    (= num  90 ) (is-prime?  90 )
    (= num  91 ) (is-prime?  91 )
    (= num  92 ) (is-prime?  92 )
    (= num  93 ) (is-prime?  93 )
    (= num  94 ) (is-prime?  94 )
    (= num  95 ) (is-prime?  95 )
    (= num  96 ) (is-prime?  96 )
    (= num  97 ) (is-prime?  97 )
    (= num  98 ) (is-prime?  98 )
    (= num  99 ) (is-prime?  99 )
    (= num  100 ) (is-prime?  100 )
    (= num  101 ) (is-prime?  101 )
    (= num  102 ) (is-prime?  102 )
    (= num  103 ) (is-prime?  103 )
    (= num  104 ) (is-prime?  104 )
    (= num  105 ) (is-prime?  105 )
    (= num  106 ) (is-prime?  106 )
    (= num  107 ) (is-prime?  107 )
    (= num  108 ) (is-prime?  108 )
    (= num  109 ) (is-prime?  109 )
    (= num  110 ) (is-prime?  110 )
    (= num  111 ) (is-prime?  111 )
    (= num  112 ) (is-prime?  112 )
    (= num  113 ) (is-prime?  113 )
    (= num  114 ) (is-prime?  114 )
    (= num  115 ) (is-prime?  115 )
    (= num  116 ) (is-prime?  116 )
    (= num  117 ) (is-prime?  117 )
    (= num  118 ) (is-prime?  118 )
    (= num  119 ) (is-prime?  119 )
    (= num  120 ) (is-prime?  120 )
    (= num  121 ) (is-prime?  121 )
    (= num  122 ) (is-prime?  122 )
    (= num  123 ) (is-prime?  123 )
    (= num  124 ) (is-prime?  124 )
    (= num  125 ) (is-prime?  125 )
    (= num  126 ) (is-prime?  126 )
    (= num  127 ) (is-prime?  127 )
    (= num  128 ) (is-prime?  128 )
    (= num  129 ) (is-prime?  129 )
    (= num  130 ) (is-prime?  130 )
    (= num  131 ) (is-prime?  131 )
    (= num  132 ) (is-prime?  132 )
    (= num  133 ) (is-prime?  133 )
    (= num  134 ) (is-prime?  134 )
    (= num  135 ) (is-prime?  135 )
    (= num  136 ) (is-prime?  136 )
    (= num  137 ) (is-prime?  137 )
    (= num  138 ) (is-prime?  138 )
    (= num  139 ) (is-prime?  139 )
    (= num  140 ) (is-prime?  140 )
    (= num  141 ) (is-prime?  141 )
    (= num  142 ) (is-prime?  142 )
    (= num  143 ) (is-prime?  143 )
    (= num  144 ) (is-prime?  144 )
    (= num  145 ) (is-prime?  145 )
    (= num  146 ) (is-prime?  146 )
    (= num  147 ) (is-prime?  147 )
    (= num  148 ) (is-prime?  148 )
    (= num  149 ) (is-prime?  149 )
    (= num  150 ) (is-prime?  150 )
    (= num  151 ) (is-prime?  151 )
    (= num  152 ) (is-prime?  152 )
    (= num  153 ) (is-prime?  153 )
    (= num  154 ) (is-prime?  154 )
    (= num  155 ) (is-prime?  155 )
    (= num  156 ) (is-prime?  156 )
    (= num  157 ) (is-prime?  157 )
    (= num  158 ) (is-prime?  158 )
    (= num  159 ) (is-prime?  159 )
    (= num  160 ) (is-prime?  160 )
    (= num  161 ) (is-prime?  161 )
    (= num  162 ) (is-prime?  162 )
    (= num  163 ) (is-prime?  163 )
    (= num  164 ) (is-prime?  164 )
    (= num  165 ) (is-prime?  165 )
    (= num  166 ) (is-prime?  166 )
    (= num  167 ) (is-prime?  167 )
    (= num  168 ) (is-prime?  168 )
    (= num  169 ) (is-prime?  169 )
    (= num  170 ) (is-prime?  170 )
    (= num  171 ) (is-prime?  171 )
    (= num  172 ) (is-prime?  172 )
    (= num  173 ) (is-prime?  173 )
    (= num  174 ) (is-prime?  174 )
    (= num  175 ) (is-prime?  175 )
    (= num  176 ) (is-prime?  176 )
    (= num  177 ) (is-prime?  177 )
    (= num  178 ) (is-prime?  178 )
    (= num  179 ) (is-prime?  179 )
    (= num  180 ) (is-prime?  180 )
    (= num  181 ) (is-prime?  181 )
    (= num  182 ) (is-prime?  182 )
    (= num  183 ) (is-prime?  183 )
    (= num  184 ) (is-prime?  184 )
    (= num  185 ) (is-prime?  185 )
    (= num  186 ) (is-prime?  186 )
    (= num  187 ) (is-prime?  187 )
    (= num  188 ) (is-prime?  188 )
    (= num  189 ) (is-prime?  189 )
    (= num  190 ) (is-prime?  190 )
    (= num  191 ) (is-prime?  191 )
    (= num  192 ) (is-prime?  192 )
    (= num  193 ) (is-prime?  193 )
    (= num  194 ) (is-prime?  194 )
    (= num  195 ) (is-prime?  195 )
    (= num  196 ) (is-prime?  196 )
    (= num  197 ) (is-prime?  197 )
    (= num  198 ) (is-prime?  198 )
    (= num  199 ) (is-prime?  199 )
    (= num  200 ) (is-prime?  200 )
    (= num  201 ) (is-prime?  201 )
    (= num  202 ) (is-prime?  202 )
    (= num  203 ) (is-prime?  203 )
    (= num  204 ) (is-prime?  204 )
    (= num  205 ) (is-prime?  205 )
    (= num  206 ) (is-prime?  206 )
    (= num  207 ) (is-prime?  207 )
    (= num  208 ) (is-prime?  208 )
    (= num  209 ) (is-prime?  209 )
    (= num  210 ) (is-prime?  210 )
    (= num  211 ) (is-prime?  211 )
    (= num  212 ) (is-prime?  212 )
    (= num  213 ) (is-prime?  213 )
    (= num  214 ) (is-prime?  214 )
    (= num  215 ) (is-prime?  215 )
    (= num  216 ) (is-prime?  216 )
    (= num  217 ) (is-prime?  217 )
    (= num  218 ) (is-prime?  218 )
    (= num  219 ) (is-prime?  219 )
    (= num  220 ) (is-prime?  220 )
    (= num  221 ) (is-prime?  221 )
    (= num  222 ) (is-prime?  222 )
    (= num  223 ) (is-prime?  223 )
    (= num  224 ) (is-prime?  224 )
    (= num  225 ) (is-prime?  225 )
    (= num  226 ) (is-prime?  226 )
    (= num  227 ) (is-prime?  227 )
    (= num  228 ) (is-prime?  228 )
    (= num  229 ) (is-prime?  229 )
    (= num  230 ) (is-prime?  230 )
    (= num  231 ) (is-prime?  231 )
    (= num  232 ) (is-prime?  232 )
    (= num  233 ) (is-prime?  233 )
    (= num  234 ) (is-prime?  234 )
    (= num  235 ) (is-prime?  235 )
    (= num  236 ) (is-prime?  236 )
    (= num  237 ) (is-prime?  237 )
    (= num  238 ) (is-prime?  238 )
    (= num  239 ) (is-prime?  239 )
    (= num  240 ) (is-prime?  240 )
    (= num  241 ) (is-prime?  241 )
    (= num  242 ) (is-prime?  242 )
    (= num  243 ) (is-prime?  243 )
    (= num  244 ) (is-prime?  244 )
    (= num  245 ) (is-prime?  245 )
    (= num  246 ) (is-prime?  246 )
    (= num  247 ) (is-prime?  247 )
    (= num  248 ) (is-prime?  248 )
    (= num  249 ) (is-prime?  249 )
    (= num  250 ) (is-prime?  250 )
    (= num  251 ) (is-prime?  251 )
    (= num  252 ) (is-prime?  252 )
    (= num  253 ) (is-prime?  253 )
    (= num  254 ) (is-prime?  254 )
    (= num  255 ) (is-prime?  255 )
    (= num  256 ) (is-prime?  256 )
    (= num  257 ) (is-prime?  257 )
    (= num  258 ) (is-prime?  258 )
    (= num  259 ) (is-prime?  259 )
    (= num  260 ) (is-prime?  260 )
    (= num  261 ) (is-prime?  261 )
    (= num  262 ) (is-prime?  262 )
    (= num  263 ) (is-prime?  263 )
    (= num  264 ) (is-prime?  264 )
    (= num  265 ) (is-prime?  265 )
    (= num  266 ) (is-prime?  266 )
    (= num  267 ) (is-prime?  267 )
    (= num  268 ) (is-prime?  268 )
    (= num  269 ) (is-prime?  269 )
    (= num  270 ) (is-prime?  270 )
    (= num  271 ) (is-prime?  271 )
    (= num  272 ) (is-prime?  272 )
    (= num  273 ) (is-prime?  273 )
    (= num  274 ) (is-prime?  274 )
    (= num  275 ) (is-prime?  275 )
    (= num  276 ) (is-prime?  276 )
    (= num  277 ) (is-prime?  277 )
    (= num  278 ) (is-prime?  278 )
    (= num  279 ) (is-prime?  279 )
    (= num  280 ) (is-prime?  280 )
    (= num  281 ) (is-prime?  281 )
    (= num  282 ) (is-prime?  282 )
    (= num  283 ) (is-prime?  283 )
    (= num  284 ) (is-prime?  284 )
    (= num  285 ) (is-prime?  285 )
    (= num  286 ) (is-prime?  286 )
    (= num  287 ) (is-prime?  287 )
    (= num  288 ) (is-prime?  288 )
    (= num  289 ) (is-prime?  289 )
    (= num  290 ) (is-prime?  290 )
    (= num  291 ) (is-prime?  291 )
    (= num  292 ) (is-prime?  292 )
    (= num  293 ) (is-prime?  293 )
    (= num  294 ) (is-prime?  294 )
    (= num  295 ) (is-prime?  295 )
    (= num  296 ) (is-prime?  296 )
    (= num  297 ) (is-prime?  297 )
    (= num  298 ) (is-prime?  298 )
    (= num  299 ) (is-prime?  299 )
    (= num  300 ) (is-prime?  300 )
    (= num  301 ) (is-prime?  301 )
    (= num  302 ) (is-prime?  302 )
    (= num  303 ) (is-prime?  303 )
    (= num  304 ) (is-prime?  304 )
    (= num  305 ) (is-prime?  305 )
    (= num  306 ) (is-prime?  306 )
    (= num  307 ) (is-prime?  307 )
    (= num  308 ) (is-prime?  308 )
    (= num  309 ) (is-prime?  309 )
    (= num  310 ) (is-prime?  310 )
    (= num  311 ) (is-prime?  311 )
    (= num  312 ) (is-prime?  312 )
    (= num  313 ) (is-prime?  313 )
    (= num  314 ) (is-prime?  314 )
    (= num  315 ) (is-prime?  315 )
    (= num  316 ) (is-prime?  316 )
    (= num  317 ) (is-prime?  317 )
    (= num  318 ) (is-prime?  318 )
    (= num  319 ) (is-prime?  319 )
    (= num  320 ) (is-prime?  320 )
    (= num  321 ) (is-prime?  321 )
    (= num  322 ) (is-prime?  322 )
    (= num  323 ) (is-prime?  323 )
    (= num  324 ) (is-prime?  324 )
    (= num  325 ) (is-prime?  325 )
    (= num  326 ) (is-prime?  326 )
    (= num  327 ) (is-prime?  327 )
    (= num  328 ) (is-prime?  328 )
    (= num  329 ) (is-prime?  329 )
    (= num  330 ) (is-prime?  330 )
    (= num  331 ) (is-prime?  331 )
    (= num  332 ) (is-prime?  332 )
    (= num  333 ) (is-prime?  333 )
    (= num  334 ) (is-prime?  334 )
    (= num  335 ) (is-prime?  335 )
    (= num  336 ) (is-prime?  336 )
    (= num  337 ) (is-prime?  337 )
    (= num  338 ) (is-prime?  338 )
    (= num  339 ) (is-prime?  339 )
    (= num  340 ) (is-prime?  340 )
    (= num  341 ) (is-prime?  341 )
    (= num  342 ) (is-prime?  342 )
    (= num  343 ) (is-prime?  343 )
    (= num  344 ) (is-prime?  344 )
    (= num  345 ) (is-prime?  345 )
    (= num  346 ) (is-prime?  346 )
    (= num  347 ) (is-prime?  347 )
    (= num  348 ) (is-prime?  348 )
    (= num  349 ) (is-prime?  349 )
    (= num  350 ) (is-prime?  350 )
    (= num  351 ) (is-prime?  351 )
    (= num  352 ) (is-prime?  352 )
    (= num  353 ) (is-prime?  353 )
    (= num  354 ) (is-prime?  354 )
    (= num  355 ) (is-prime?  355 )
    (= num  356 ) (is-prime?  356 )
    (= num  357 ) (is-prime?  357 )
    (= num  358 ) (is-prime?  358 )
    (= num  359 ) (is-prime?  359 )
    (= num  360 ) (is-prime?  360 )
    (= num  361 ) (is-prime?  361 )
    (= num  362 ) (is-prime?  362 )
    (= num  363 ) (is-prime?  363 )
    (= num  364 ) (is-prime?  364 )
    (= num  365 ) (is-prime?  365 )
    (= num  366 ) (is-prime?  366 )
    (= num  367 ) (is-prime?  367 )
    (= num  368 ) (is-prime?  368 )
    (= num  369 ) (is-prime?  369 )
    (= num  370 ) (is-prime?  370 )
    (= num  371 ) (is-prime?  371 )
    (= num  372 ) (is-prime?  372 )
    (= num  373 ) (is-prime?  373 )
    (= num  374 ) (is-prime?  374 )
    (= num  375 ) (is-prime?  375 )
    (= num  376 ) (is-prime?  376 )
    (= num  377 ) (is-prime?  377 )
    (= num  378 ) (is-prime?  378 )
    (= num  379 ) (is-prime?  379 )
    (= num  380 ) (is-prime?  380 )
    (= num  381 ) (is-prime?  381 )
    (= num  382 ) (is-prime?  382 )
    (= num  383 ) (is-prime?  383 )
    (= num  384 ) (is-prime?  384 )
    (= num  385 ) (is-prime?  385 )
    (= num  386 ) (is-prime?  386 )
    (= num  387 ) (is-prime?  387 )
    (= num  388 ) (is-prime?  388 )
    (= num  389 ) (is-prime?  389 )
    (= num  390 ) (is-prime?  390 )
    (= num  391 ) (is-prime?  391 )
    (= num  392 ) (is-prime?  392 )
    (= num  393 ) (is-prime?  393 )
    (= num  394 ) (is-prime?  394 )
    (= num  395 ) (is-prime?  395 )
    (= num  396 ) (is-prime?  396 )
    (= num  397 ) (is-prime?  397 )
    (= num  398 ) (is-prime?  398 )
    (= num  399 ) (is-prime?  399 )
    (= num  400 ) (is-prime?  400 )
    (= num  401 ) (is-prime?  401 )
    (= num  402 ) (is-prime?  402 )
    (= num  403 ) (is-prime?  403 )
    (= num  404 ) (is-prime?  404 )
    (= num  405 ) (is-prime?  405 )
    (= num  406 ) (is-prime?  406 )
    (= num  407 ) (is-prime?  407 )
    (= num  408 ) (is-prime?  408 )
    (= num  409 ) (is-prime?  409 )
    (= num  410 ) (is-prime?  410 )
    (= num  411 ) (is-prime?  411 )
    (= num  412 ) (is-prime?  412 )
    (= num  413 ) (is-prime?  413 )
    (= num  414 ) (is-prime?  414 )
    (= num  415 ) (is-prime?  415 )
    (= num  416 ) (is-prime?  416 )
    (= num  417 ) (is-prime?  417 )
    (= num  418 ) (is-prime?  418 )
    (= num  419 ) (is-prime?  419 )
    (= num  420 ) (is-prime?  420 )
    (= num  421 ) (is-prime?  421 )
    (= num  422 ) (is-prime?  422 )
    (= num  423 ) (is-prime?  423 )
    (= num  424 ) (is-prime?  424 )
    (= num  425 ) (is-prime?  425 )
    (= num  426 ) (is-prime?  426 )
    (= num  427 ) (is-prime?  427 )
    (= num  428 ) (is-prime?  428 )
    (= num  429 ) (is-prime?  429 )
    (= num  430 ) (is-prime?  430 )
    (= num  431 ) (is-prime?  431 )
    (= num  432 ) (is-prime?  432 )
    (= num  433 ) (is-prime?  433 )
    (= num  434 ) (is-prime?  434 )
    (= num  435 ) (is-prime?  435 )
    (= num  436 ) (is-prime?  436 )
    (= num  437 ) (is-prime?  437 )
    (= num  438 ) (is-prime?  438 )
    (= num  439 ) (is-prime?  439 )
    (= num  440 ) (is-prime?  440 )
    (= num  441 ) (is-prime?  441 )
    (= num  442 ) (is-prime?  442 )
    (= num  443 ) (is-prime?  443 )
    (= num  444 ) (is-prime?  444 )
    (= num  445 ) (is-prime?  445 )
    (= num  446 ) (is-prime?  446 )
    (= num  447 ) (is-prime?  447 )
    (= num  448 ) (is-prime?  448 )
    (= num  449 ) (is-prime?  449 )
    (= num  450 ) (is-prime?  450 )
    (= num  451 ) (is-prime?  451 )
    (= num  452 ) (is-prime?  452 )
    (= num  453 ) (is-prime?  453 )
    (= num  454 ) (is-prime?  454 )
    (= num  455 ) (is-prime?  455 )
    (= num  456 ) (is-prime?  456 )
    (= num  457 ) (is-prime?  457 )
    (= num  458 ) (is-prime?  458 )
    (= num  459 ) (is-prime?  459 )
    (= num  460 ) (is-prime?  460 )
    (= num  461 ) (is-prime?  461 )
    (= num  462 ) (is-prime?  462 )
    (= num  463 ) (is-prime?  463 )
    (= num  464 ) (is-prime?  464 )
    (= num  465 ) (is-prime?  465 )
    (= num  466 ) (is-prime?  466 )
    (= num  467 ) (is-prime?  467 )
    (= num  468 ) (is-prime?  468 )
    (= num  469 ) (is-prime?  469 )
    (= num  470 ) (is-prime?  470 )
    (= num  471 ) (is-prime?  471 )
    (= num  472 ) (is-prime?  472 )
    (= num  473 ) (is-prime?  473 )
    (= num  474 ) (is-prime?  474 )
    (= num  475 ) (is-prime?  475 )
    (= num  476 ) (is-prime?  476 )
    (= num  477 ) (is-prime?  477 )
    (= num  478 ) (is-prime?  478 )
    (= num  479 ) (is-prime?  479 )
    (= num  480 ) (is-prime?  480 )
    (= num  481 ) (is-prime?  481 )
    (= num  482 ) (is-prime?  482 )
    (= num  483 ) (is-prime?  483 )
    (= num  484 ) (is-prime?  484 )
    (= num  485 ) (is-prime?  485 )
    (= num  486 ) (is-prime?  486 )
    (= num  487 ) (is-prime?  487 )
    (= num  488 ) (is-prime?  488 )
    (= num  489 ) (is-prime?  489 )
    (= num  490 ) (is-prime?  490 )
    (= num  491 ) (is-prime?  491 )
    (= num  492 ) (is-prime?  492 )
    (= num  493 ) (is-prime?  493 )
    (= num  494 ) (is-prime?  494 )
    (= num  495 ) (is-prime?  495 )
    (= num  496 ) (is-prime?  496 )
    (= num  497 ) (is-prime?  497 )
    (= num  498 ) (is-prime?  498 )
    (= num  499 ) (is-prime?  499 )
    (= num  500 ) (is-prime?  500 )
    (= num  501 ) (is-prime?  501 )
    (= num  502 ) (is-prime?  502 )
    (= num  503 ) (is-prime?  503 )
    (= num  504 ) (is-prime?  504 )
    (= num  505 ) (is-prime?  505 )
    (= num  506 ) (is-prime?  506 )
    (= num  507 ) (is-prime?  507 )
    (= num  508 ) (is-prime?  508 )
    (= num  509 ) (is-prime?  509 )
    (= num  510 ) (is-prime?  510 )
    (= num  511 ) (is-prime?  511 )
    (= num  512 ) (is-prime?  512 )
    (= num  513 ) (is-prime?  513 )
    (= num  514 ) (is-prime?  514 )
    (= num  515 ) (is-prime?  515 )
    (= num  516 ) (is-prime?  516 )
    (= num  517 ) (is-prime?  517 )
    (= num  518 ) (is-prime?  518 )
    (= num  519 ) (is-prime?  519 )
    (= num  520 ) (is-prime?  520 )
    (= num  521 ) (is-prime?  521 )
    (= num  522 ) (is-prime?  522 )
    (= num  523 ) (is-prime?  523 )
    (= num  524 ) (is-prime?  524 )
    (= num  525 ) (is-prime?  525 )
    (= num  526 ) (is-prime?  526 )
    (= num  527 ) (is-prime?  527 )
    (= num  528 ) (is-prime?  528 )
    (= num  529 ) (is-prime?  529 )
    (= num  530 ) (is-prime?  530 )
    (= num  531 ) (is-prime?  531 )
    (= num  532 ) (is-prime?  532 )
    (= num  533 ) (is-prime?  533 )
    (= num  534 ) (is-prime?  534 )
    (= num  535 ) (is-prime?  535 )
    (= num  536 ) (is-prime?  536 )
    (= num  537 ) (is-prime?  537 )
    (= num  538 ) (is-prime?  538 )
    (= num  539 ) (is-prime?  539 )
    (= num  540 ) (is-prime?  540 )
    (= num  541 ) (is-prime?  541 )
    (= num  542 ) (is-prime?  542 )
    (= num  543 ) (is-prime?  543 )
    (= num  544 ) (is-prime?  544 )
    (= num  545 ) (is-prime?  545 )
    (= num  546 ) (is-prime?  546 )
    (= num  547 ) (is-prime?  547 )
    (= num  548 ) (is-prime?  548 )
    (= num  549 ) (is-prime?  549 )
    (= num  550 ) (is-prime?  550 )
    (= num  551 ) (is-prime?  551 )
    (= num  552 ) (is-prime?  552 )
    (= num  553 ) (is-prime?  553 )
    (= num  554 ) (is-prime?  554 )
    (= num  555 ) (is-prime?  555 )
    (= num  556 ) (is-prime?  556 )
    (= num  557 ) (is-prime?  557 )
    (= num  558 ) (is-prime?  558 )
    (= num  559 ) (is-prime?  559 )
    (= num  560 ) (is-prime?  560 )
    (= num  561 ) (is-prime?  561 )
    (= num  562 ) (is-prime?  562 )
    (= num  563 ) (is-prime?  563 )
    (= num  564 ) (is-prime?  564 )
    (= num  565 ) (is-prime?  565 )
    (= num  566 ) (is-prime?  566 )
    (= num  567 ) (is-prime?  567 )
    (= num  568 ) (is-prime?  568 )
    (= num  569 ) (is-prime?  569 )
    (= num  570 ) (is-prime?  570 )
    (= num  571 ) (is-prime?  571 )
    (= num  572 ) (is-prime?  572 )
    (= num  573 ) (is-prime?  573 )
    (= num  574 ) (is-prime?  574 )
    (= num  575 ) (is-prime?  575 )
    (= num  576 ) (is-prime?  576 )
    (= num  577 ) (is-prime?  577 )
    (= num  578 ) (is-prime?  578 )
    (= num  579 ) (is-prime?  579 )
    (= num  580 ) (is-prime?  580 )
    (= num  581 ) (is-prime?  581 )
    (= num  582 ) (is-prime?  582 )
    (= num  583 ) (is-prime?  583 )
    (= num  584 ) (is-prime?  584 )
    (= num  585 ) (is-prime?  585 )
    (= num  586 ) (is-prime?  586 )
    (= num  587 ) (is-prime?  587 )
    (= num  588 ) (is-prime?  588 )
    (= num  589 ) (is-prime?  589 )
    (= num  590 ) (is-prime?  590 )
    (= num  591 ) (is-prime?  591 )
    (= num  592 ) (is-prime?  592 )
    (= num  593 ) (is-prime?  593 )
    (= num  594 ) (is-prime?  594 )
    (= num  595 ) (is-prime?  595 )
    (= num  596 ) (is-prime?  596 )
    (= num  597 ) (is-prime?  597 )
    (= num  598 ) (is-prime?  598 )
    (= num  599 ) (is-prime?  599 )
    (= num  600 ) (is-prime?  600 )
    (= num  601 ) (is-prime?  601 )
    (= num  602 ) (is-prime?  602 )
    (= num  603 ) (is-prime?  603 )
    (= num  604 ) (is-prime?  604 )
    (= num  605 ) (is-prime?  605 )
    (= num  606 ) (is-prime?  606 )
    (= num  607 ) (is-prime?  607 )
    (= num  608 ) (is-prime?  608 )
    (= num  609 ) (is-prime?  609 )
    (= num  610 ) (is-prime?  610 )
    (= num  611 ) (is-prime?  611 )
    (= num  612 ) (is-prime?  612 )
    (= num  613 ) (is-prime?  613 )
    (= num  614 ) (is-prime?  614 )
    (= num  615 ) (is-prime?  615 )
    (= num  616 ) (is-prime?  616 )
    (= num  617 ) (is-prime?  617 )
    (= num  618 ) (is-prime?  618 )
    (= num  619 ) (is-prime?  619 )
    (= num  620 ) (is-prime?  620 )
    (= num  621 ) (is-prime?  621 )
    (= num  622 ) (is-prime?  622 )
    (= num  623 ) (is-prime?  623 )
    (= num  624 ) (is-prime?  624 )
    (= num  625 ) (is-prime?  625 )
    (= num  626 ) (is-prime?  626 )
    (= num  627 ) (is-prime?  627 )
    (= num  628 ) (is-prime?  628 )
    (= num  629 ) (is-prime?  629 )
    (= num  630 ) (is-prime?  630 )
    (= num  631 ) (is-prime?  631 )
    (= num  632 ) (is-prime?  632 )
    (= num  633 ) (is-prime?  633 )
    (= num  634 ) (is-prime?  634 )
    (= num  635 ) (is-prime?  635 )
    (= num  636 ) (is-prime?  636 )
    (= num  637 ) (is-prime?  637 )
    (= num  638 ) (is-prime?  638 )
    (= num  639 ) (is-prime?  639 )
    (= num  640 ) (is-prime?  640 )
    (= num  641 ) (is-prime?  641 )
    (= num  642 ) (is-prime?  642 )
    (= num  643 ) (is-prime?  643 )
    (= num  644 ) (is-prime?  644 )
    (= num  645 ) (is-prime?  645 )
    (= num  646 ) (is-prime?  646 )
    (= num  647 ) (is-prime?  647 )
    (= num  648 ) (is-prime?  648 )
    (= num  649 ) (is-prime?  649 )
    (= num  650 ) (is-prime?  650 )
    (= num  651 ) (is-prime?  651 )
    (= num  652 ) (is-prime?  652 )
    (= num  653 ) (is-prime?  653 )
    (= num  654 ) (is-prime?  654 )
    (= num  655 ) (is-prime?  655 )
    (= num  656 ) (is-prime?  656 )
    (= num  657 ) (is-prime?  657 )
    (= num  658 ) (is-prime?  658 )
    (= num  659 ) (is-prime?  659 )
    (= num  660 ) (is-prime?  660 )
    (= num  661 ) (is-prime?  661 )
    (= num  662 ) (is-prime?  662 )
    (= num  663 ) (is-prime?  663 )
    (= num  664 ) (is-prime?  664 )
    (= num  665 ) (is-prime?  665 )
    (= num  666 ) (is-prime?  666 )
    (= num  667 ) (is-prime?  667 )
    (= num  668 ) (is-prime?  668 )
    (= num  669 ) (is-prime?  669 )
    (= num  670 ) (is-prime?  670 )
    (= num  671 ) (is-prime?  671 )
    (= num  672 ) (is-prime?  672 )
    (= num  673 ) (is-prime?  673 )
    (= num  674 ) (is-prime?  674 )
    (= num  675 ) (is-prime?  675 )
    (= num  676 ) (is-prime?  676 )
    (= num  677 ) (is-prime?  677 )
    (= num  678 ) (is-prime?  678 )
    (= num  679 ) (is-prime?  679 )
    (= num  680 ) (is-prime?  680 )
    (= num  681 ) (is-prime?  681 )
    (= num  682 ) (is-prime?  682 )
    (= num  683 ) (is-prime?  683 )
    (= num  684 ) (is-prime?  684 )
    (= num  685 ) (is-prime?  685 )
    (= num  686 ) (is-prime?  686 )
    (= num  687 ) (is-prime?  687 )
    (= num  688 ) (is-prime?  688 )
    (= num  689 ) (is-prime?  689 )
    (= num  690 ) (is-prime?  690 )
    (= num  691 ) (is-prime?  691 )
    (= num  692 ) (is-prime?  692 )
    (= num  693 ) (is-prime?  693 )
    (= num  694 ) (is-prime?  694 )
    (= num  695 ) (is-prime?  695 )
    (= num  696 ) (is-prime?  696 )
    (= num  697 ) (is-prime?  697 )
    (= num  698 ) (is-prime?  698 )
    (= num  699 ) (is-prime?  699 )
    (= num  700 ) (is-prime?  700 )
    (= num  701 ) (is-prime?  701 )
    (= num  702 ) (is-prime?  702 )
    (= num  703 ) (is-prime?  703 )
    (= num  704 ) (is-prime?  704 )
    (= num  705 ) (is-prime?  705 )
    (= num  706 ) (is-prime?  706 )
    (= num  707 ) (is-prime?  707 )
    (= num  708 ) (is-prime?  708 )
    (= num  709 ) (is-prime?  709 )
    (= num  710 ) (is-prime?  710 )
    (= num  711 ) (is-prime?  711 )
    (= num  712 ) (is-prime?  712 )
    (= num  713 ) (is-prime?  713 )
    (= num  714 ) (is-prime?  714 )
    (= num  715 ) (is-prime?  715 )
    (= num  716 ) (is-prime?  716 )
    (= num  717 ) (is-prime?  717 )
    (= num  718 ) (is-prime?  718 )
    (= num  719 ) (is-prime?  719 )
    (= num  720 ) (is-prime?  720 )
    (= num  721 ) (is-prime?  721 )
    (= num  722 ) (is-prime?  722 )
    (= num  723 ) (is-prime?  723 )
    (= num  724 ) (is-prime?  724 )
    (= num  725 ) (is-prime?  725 )
    (= num  726 ) (is-prime?  726 )
    (= num  727 ) (is-prime?  727 )
    (= num  728 ) (is-prime?  728 )
    (= num  729 ) (is-prime?  729 )
    (= num  730 ) (is-prime?  730 )
    (= num  731 ) (is-prime?  731 )
    (= num  732 ) (is-prime?  732 )
    (= num  733 ) (is-prime?  733 )
    (= num  734 ) (is-prime?  734 )
    (= num  735 ) (is-prime?  735 )
    (= num  736 ) (is-prime?  736 )
    (= num  737 ) (is-prime?  737 )
    (= num  738 ) (is-prime?  738 )
    (= num  739 ) (is-prime?  739 )
    (= num  740 ) (is-prime?  740 )
    (= num  741 ) (is-prime?  741 )
    (= num  742 ) (is-prime?  742 )
    (= num  743 ) (is-prime?  743 )
    (= num  744 ) (is-prime?  744 )
    (= num  745 ) (is-prime?  745 )
    (= num  746 ) (is-prime?  746 )
    (= num  747 ) (is-prime?  747 )
    (= num  748 ) (is-prime?  748 )
    (= num  749 ) (is-prime?  749 )
    (= num  750 ) (is-prime?  750 )
    (= num  751 ) (is-prime?  751 )
    (= num  752 ) (is-prime?  752 )
    (= num  753 ) (is-prime?  753 )
    (= num  754 ) (is-prime?  754 )
    (= num  755 ) (is-prime?  755 )
    (= num  756 ) (is-prime?  756 )
    (= num  757 ) (is-prime?  757 )
    (= num  758 ) (is-prime?  758 )
    (= num  759 ) (is-prime?  759 )
    (= num  760 ) (is-prime?  760 )
    (= num  761 ) (is-prime?  761 )
    (= num  762 ) (is-prime?  762 )
    (= num  763 ) (is-prime?  763 )
    (= num  764 ) (is-prime?  764 )
    (= num  765 ) (is-prime?  765 )
    (= num  766 ) (is-prime?  766 )
    (= num  767 ) (is-prime?  767 )
    (= num  768 ) (is-prime?  768 )
    (= num  769 ) (is-prime?  769 )
    (= num  770 ) (is-prime?  770 )
    (= num  771 ) (is-prime?  771 )
    (= num  772 ) (is-prime?  772 )
    (= num  773 ) (is-prime?  773 )
    (= num  774 ) (is-prime?  774 )
    (= num  775 ) (is-prime?  775 )
    (= num  776 ) (is-prime?  776 )
    (= num  777 ) (is-prime?  777 )
    (= num  778 ) (is-prime?  778 )
    (= num  779 ) (is-prime?  779 )
    (= num  780 ) (is-prime?  780 )
    (= num  781 ) (is-prime?  781 )
    (= num  782 ) (is-prime?  782 )
    (= num  783 ) (is-prime?  783 )
    (= num  784 ) (is-prime?  784 )
    (= num  785 ) (is-prime?  785 )
    (= num  786 ) (is-prime?  786 )
    (= num  787 ) (is-prime?  787 )
    (= num  788 ) (is-prime?  788 )
    (= num  789 ) (is-prime?  789 )
    (= num  790 ) (is-prime?  790 )
    (= num  791 ) (is-prime?  791 )
    (= num  792 ) (is-prime?  792 )
    (= num  793 ) (is-prime?  793 )
    (= num  794 ) (is-prime?  794 )
    (= num  795 ) (is-prime?  795 )
    (= num  796 ) (is-prime?  796 )
    (= num  797 ) (is-prime?  797 )
    (= num  798 ) (is-prime?  798 )
    (= num  799 ) (is-prime?  799 )
    (= num  800 ) (is-prime?  800 )
    (= num  801 ) (is-prime?  801 )
    (= num  802 ) (is-prime?  802 )
    (= num  803 ) (is-prime?  803 )
    (= num  804 ) (is-prime?  804 )
    (= num  805 ) (is-prime?  805 )
    (= num  806 ) (is-prime?  806 )
    (= num  807 ) (is-prime?  807 )
    (= num  808 ) (is-prime?  808 )
    (= num  809 ) (is-prime?  809 )
    (= num  810 ) (is-prime?  810 )
    (= num  811 ) (is-prime?  811 )
    (= num  812 ) (is-prime?  812 )
    (= num  813 ) (is-prime?  813 )
    (= num  814 ) (is-prime?  814 )
    (= num  815 ) (is-prime?  815 )
    (= num  816 ) (is-prime?  816 )
    (= num  817 ) (is-prime?  817 )
    (= num  818 ) (is-prime?  818 )
    (= num  819 ) (is-prime?  819 )
    (= num  820 ) (is-prime?  820 )
    (= num  821 ) (is-prime?  821 )
    (= num  822 ) (is-prime?  822 )
    (= num  823 ) (is-prime?  823 )
    (= num  824 ) (is-prime?  824 )
    (= num  825 ) (is-prime?  825 )
    (= num  826 ) (is-prime?  826 )
    (= num  827 ) (is-prime?  827 )
    (= num  828 ) (is-prime?  828 )
    (= num  829 ) (is-prime?  829 )
    (= num  830 ) (is-prime?  830 )
    (= num  831 ) (is-prime?  831 )
    (= num  832 ) (is-prime?  832 )
    (= num  833 ) (is-prime?  833 )
    (= num  834 ) (is-prime?  834 )
    (= num  835 ) (is-prime?  835 )
    (= num  836 ) (is-prime?  836 )
    (= num  837 ) (is-prime?  837 )
    (= num  838 ) (is-prime?  838 )
    (= num  839 ) (is-prime?  839 )
    (= num  840 ) (is-prime?  840 )
    (= num  841 ) (is-prime?  841 )
    (= num  842 ) (is-prime?  842 )
    (= num  843 ) (is-prime?  843 )
    (= num  844 ) (is-prime?  844 )
    (= num  845 ) (is-prime?  845 )
    (= num  846 ) (is-prime?  846 )
    (= num  847 ) (is-prime?  847 )
    (= num  848 ) (is-prime?  848 )
    (= num  849 ) (is-prime?  849 )
    (= num  850 ) (is-prime?  850 )
    (= num  851 ) (is-prime?  851 )
    (= num  852 ) (is-prime?  852 )
    (= num  853 ) (is-prime?  853 )
    (= num  854 ) (is-prime?  854 )
    (= num  855 ) (is-prime?  855 )
    (= num  856 ) (is-prime?  856 )
    (= num  857 ) (is-prime?  857 )
    (= num  858 ) (is-prime?  858 )
    (= num  859 ) (is-prime?  859 )
    (= num  860 ) (is-prime?  860 )
    (= num  861 ) (is-prime?  861 )
    (= num  862 ) (is-prime?  862 )
    (= num  863 ) (is-prime?  863 )
    (= num  864 ) (is-prime?  864 )
    (= num  865 ) (is-prime?  865 )
    (= num  866 ) (is-prime?  866 )
    (= num  867 ) (is-prime?  867 )
    (= num  868 ) (is-prime?  868 )
    (= num  869 ) (is-prime?  869 )
    (= num  870 ) (is-prime?  870 )
    (= num  871 ) (is-prime?  871 )
    (= num  872 ) (is-prime?  872 )
    (= num  873 ) (is-prime?  873 )
    (= num  874 ) (is-prime?  874 )
    (= num  875 ) (is-prime?  875 )
    (= num  876 ) (is-prime?  876 )
    (= num  877 ) (is-prime?  877 )
    (= num  878 ) (is-prime?  878 )
    (= num  879 ) (is-prime?  879 )
    (= num  880 ) (is-prime?  880 )
    (= num  881 ) (is-prime?  881 )
    (= num  882 ) (is-prime?  882 )
    (= num  883 ) (is-prime?  883 )
    (= num  884 ) (is-prime?  884 )
    (= num  885 ) (is-prime?  885 )
    (= num  886 ) (is-prime?  886 )
    (= num  887 ) (is-prime?  887 )
    (= num  888 ) (is-prime?  888 )
    (= num  889 ) (is-prime?  889 )
    (= num  890 ) (is-prime?  890 )
    (= num  891 ) (is-prime?  891 )
    (= num  892 ) (is-prime?  892 )
    (= num  893 ) (is-prime?  893 )
    (= num  894 ) (is-prime?  894 )
    (= num  895 ) (is-prime?  895 )
    (= num  896 ) (is-prime?  896 )
    (= num  897 ) (is-prime?  897 )
    (= num  898 ) (is-prime?  898 )
    (= num  899 ) (is-prime?  899 )
    (= num  900 ) (is-prime?  900 )
    (= num  901 ) (is-prime?  901 )
    (= num  902 ) (is-prime?  902 )
    (= num  903 ) (is-prime?  903 )
    (= num  904 ) (is-prime?  904 )
    (= num  905 ) (is-prime?  905 )
    (= num  906 ) (is-prime?  906 )
    (= num  907 ) (is-prime?  907 )
    (= num  908 ) (is-prime?  908 )
    (= num  909 ) (is-prime?  909 )
    (= num  910 ) (is-prime?  910 )
    (= num  911 ) (is-prime?  911 )
    (= num  912 ) (is-prime?  912 )
    (= num  913 ) (is-prime?  913 )
    (= num  914 ) (is-prime?  914 )
    (= num  915 ) (is-prime?  915 )
    (= num  916 ) (is-prime?  916 )
    (= num  917 ) (is-prime?  917 )
    (= num  918 ) (is-prime?  918 )
    (= num  919 ) (is-prime?  919 )
    (= num  920 ) (is-prime?  920 )
    (= num  921 ) (is-prime?  921 )
    (= num  922 ) (is-prime?  922 )
    (= num  923 ) (is-prime?  923 )
    (= num  924 ) (is-prime?  924 )
    (= num  925 ) (is-prime?  925 )
    (= num  926 ) (is-prime?  926 )
    (= num  927 ) (is-prime?  927 )
    (= num  928 ) (is-prime?  928 )
    (= num  929 ) (is-prime?  929 )
    (= num  930 ) (is-prime?  930 )
    (= num  931 ) (is-prime?  931 )
    (= num  932 ) (is-prime?  932 )
    (= num  933 ) (is-prime?  933 )
    (= num  934 ) (is-prime?  934 )
    (= num  935 ) (is-prime?  935 )
    (= num  936 ) (is-prime?  936 )
    (= num  937 ) (is-prime?  937 )
    (= num  938 ) (is-prime?  938 )
    (= num  939 ) (is-prime?  939 )
    (= num  940 ) (is-prime?  940 )
    (= num  941 ) (is-prime?  941 )
    (= num  942 ) (is-prime?  942 )
    (= num  943 ) (is-prime?  943 )
    (= num  944 ) (is-prime?  944 )
    (= num  945 ) (is-prime?  945 )
    (= num  946 ) (is-prime?  946 )
    (= num  947 ) (is-prime?  947 )
    (= num  948 ) (is-prime?  948 )
    (= num  949 ) (is-prime?  949 )
    (= num  950 ) (is-prime?  950 )
    (= num  951 ) (is-prime?  951 )
    (= num  952 ) (is-prime?  952 )
    (= num  953 ) (is-prime?  953 )
    (= num  954 ) (is-prime?  954 )
    (= num  955 ) (is-prime?  955 )
    (= num  956 ) (is-prime?  956 )
    (= num  957 ) (is-prime?  957 )
    (= num  958 ) (is-prime?  958 )
    (= num  959 ) (is-prime?  959 )
    (= num  960 ) (is-prime?  960 )
    (= num  961 ) (is-prime?  961 )
    (= num  962 ) (is-prime?  962 )
    (= num  963 ) (is-prime?  963 )
    (= num  964 ) (is-prime?  964 )
    (= num  965 ) (is-prime?  965 )
    (= num  966 ) (is-prime?  966 )
    (= num  967 ) (is-prime?  967 )
    (= num  968 ) (is-prime?  968 )
    (= num  969 ) (is-prime?  969 )
    (= num  970 ) (is-prime?  970 )
    (= num  971 ) (is-prime?  971 )
    (= num  972 ) (is-prime?  972 )
    (= num  973 ) (is-prime?  973 )
    (= num  974 ) (is-prime?  974 )
    (= num  975 ) (is-prime?  975 )
    (= num  976 ) (is-prime?  976 )
    (= num  977 ) (is-prime?  977 )
    (= num  978 ) (is-prime?  978 )
    (= num  979 ) (is-prime?  979 )
    (= num  980 ) (is-prime?  980 )
    (= num  981 ) (is-prime?  981 )
    (= num  982 ) (is-prime?  982 )
    (= num  983 ) (is-prime?  983 )
    (= num  984 ) (is-prime?  984 )
    (= num  985 ) (is-prime?  985 )
    (= num  986 ) (is-prime?  986 )
    (= num  987 ) (is-prime?  987 )
    (= num  988 ) (is-prime?  988 )
    (= num  989 ) (is-prime?  989 )
    (= num  990 ) (is-prime?  990 )
    (= num  991 ) (is-prime?  991 )
    (= num  992 ) (is-prime?  992 )
    (= num  993 ) (is-prime?  993 )
    (= num  994 ) (is-prime?  994 )
    (= num  995 ) (is-prime?  995 )
    (= num  996 ) (is-prime?  996 )
    (= num  997 ) (is-prime?  997 )
    (= num  998 ) (is-prime?  998 )
    (= num  999 ) (is-prime?  999 )
    (= num 1000) (is-prime? 1000)
    :else false))

(defn -main []
  (print "Enter a number: ")
  (flush)
  (let [num (read-line)
        num (Integer. num)]
    (if (check-prime num)
      (println (str num " is a prime number."))
      (println (str num " is not a prime number.")))))

(-main)