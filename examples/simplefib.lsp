(
 (zero ())
 (one (1))
 (and (lambda (x y)
        (cond
          (x (cond (y 't) ('t '())))
          ('t '()))))
 (adds (lambda (a b)
         (cond ((atom b) a)
               ('t (cons (car b) (adds (cdr b) a))))))
 (subs (lambda (a b)
         (cond ((atom b) a)
               ((atom a) b)
               ('t (subs (cdr a) (cdr b))))))
 (equals (lambda (a b)
           (cond ((and (atom a) (atom b)) 't)
                 ((atom a) '())
                 ((atom b) '())
                 ('t (equals (cdr a) (cdr b))))))
 (fib (lambda (a b n)
        (cond
          ((equals n zero) a)
          ('t (fib b (adds a b) (subs n one)))
          )))
 (main (lambda () (fib zero one '(1 1 1 1 1 1 1 1))))
 )


