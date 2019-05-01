(
 (adds (lambda (a b)
         (cond ((null b) a)
               ('t (cons (car b) (adds (cdr b) a))))))
 (subs (lambda (a b)
         (cond ((null b) a)
               ((null a) b)
               ('t (subs (cdr a) (cdr b))))))
 (equals (lambda (a b)
           (cond ((and (null a) (null b)) 't)
                 ((null a) '())
                 ((null b) '())
                 ('t (equals (cdr a) (cdr b))))))
)


