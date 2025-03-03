GOOF----LE-8-2.0�X      ]� 4  h�      ] g  guile�	 �	g  define-module*�	 �	 �	g  language�	g  
ecmascript�	g  tokenize�		 �	
g  filenameS�	f   language/ecmascript/tokenize.scm�	g  importsS�	g  ice-9�	g  rdelim�	 �	 �	g  srfi�	g  srfi-1�	 �	g  selectS�	g  unfold-right�	 �	 �	g  system�	g  base�	g  lalr�	 �	 �	 �	g  exportsS�	g  
next-token�	 g  make-tokenizer�	!g  make-tokenizer/1�	"g  
tokenize/1�	# !" �	$g  set-current-module�	%$ �	&$ �	'g  throw�	(g  syntax-error�	)g  and=>�	*g  "source-location->source-properties�	+g  port-filename�	,g  	port-line�	-g  port-column�	.g  catch�	/. �	0. �	1g  ftell�	2g  source-location�	32 �	42 �	5g  port-source-location�	6g  eof-object?�	7g  	peek-char�	8f  EOF while reading a token�	9g  read-delimited�	:g  peek�	;g  
read-until�	<g  char-numeric?�	=g  memv�	>abcdef �	?ABCDEF �	@g  	char-hex?�	Ag  char->integer�	Bg  digit->number�	Cg  char-downcase�	Dg  hex->number�	Eg  	read-char�	Fg  lexical-token�	GF �	HF �	Ig  /�	Jg  char=?�	Kg  	read-line�	Lf  EOF while in multi-line comment�	Mg  /=�	Ng  read-regexp�	Og  
read-slash�	Pg  string�	Qg  char-alphabetic?�	Rg  string-append�	Sg  reverse�	Tg  RegexpLiteral�	Uf  (regexp literals may not contain newlines�	Vf   �	Wg  StringLiteral�	Xf  (octal escape sequences are not supported�	Yf  bad hex character escape�	Zg  integer->char�	[g  string->number�	\f  (string literals may not contain newlines�	]g  read-string�	^f  break�	_g  break�	`^_��	af  else�	bg  else�	cab��	df  new�	eg  new�	fde��	gf  var�	hg  var�	igh��	jf  case�	kg  case�	ljk��	mf  finally�	ng  finally�	omn��	pf  return�	qg  return�	rpq��	sf  void�	tg  void�	ust��	vf  catch�	wv.��	xf  for�	yg  for�	zxy��	{f  switch�	|g  switch�	}{|��	~f  while�	g  while� �~�� �f  continue� �g  continue� ����� �f  function� �g  function� ����� �f  this� �g  this� ����� �f  with� �g  with� ����� �f  default� �g  default� ����� �f  if� �g  if� ����� �f  throw� ��'�� �f  delete� �g  delete� ����� �f  in� �g  in� ����� �f  try� �g  try� ����� �f  do� �g  do� ����� �f  
instanceof� �g  
instanceof� ����� �f  typeof� �g  typeof� ����� �f  null� �g  null� ����� �f  true� �g  true� ����� �f  false� �g  false� ����� �`cfiloruwz}����������������� � �g  
*keywords*� �f  abstract� �g  abstract� ����� �f  enum� �g  enum� ����� �f  int� �g  int� ����� �f  short� �g  short� ����� �f  boolean� �g  boolean� ����� �f  export� �g  export� ��� �f  	interface� �g  	interface� ��Ō� �f  static� �g  static� ��Ȍ� �f  byte� �g  byte� ��ˌ� �f  extends� �g  extends� ��Ό� �f  long� �g  long� ��ь� �f  super� �g  super� ��Ԍ� �f  char� �g  char� ��׌� �f  final� �g  final� ��ڌ� �f  native� �g  native� ��݌� �f  synchronized� �g  synchronized� ����� �f  class� �g  class� ��㌤ �f  float� �g  float� ��挤 �f  package� �g  package� ��錤 �f  throws� �g  throws� ��쌤 �f  const� �g  const� �� �f  goto� �g  goto� ��� �f  private� �g  private� ����� �f  	transient� �g  	transient� ����� �f  debugger� �g  debugger� ����� �f  
implements� �g  
implements� ����� f  	protected�g  	protected� ��f  volatile�g  volatile���f  double�g  double���	f  import�
g  import�	
��f  public�g  public����������������������������� �g  *future-reserved-words*�g  list->string�g  	assoc-ref�f  &word is reserved for the future, dude.�g  string->symbol�g  
Identifier�g  read-identifier�f  error reading exponent: EOF�g  +�g  -�f  !error reading exponent: non-digit�e  1.0�g  expt�e  0.0�f  $bad digit reading hexadecimal number�f  invalid digit in octal sequence� g  read-numeric�!f  {�"g  lbrace�#!"��$f  }�%g  rbrace�&$%��'f  (�(g  lparen�)'(��*f  )�+g  rparen�,*+��-f  [�.g  lbracket�/-.��0f  ]�1g  rbracket�201��3f  .�4g  dot�534��6f  ;�7g  	semicolon�867��9f  ,�:g  comma�;9:��<f  <�=g  <�><=��?f  >�@g  >�A?@��Bf  <=�Cg  <=�DBC��Ef  >=�Fg  >=�GEF��Hf  ==�Ig  ==�JHI��Kf  !=�Lg  !=�MKL��Nf  ===�Og  ===�PNO��Qf  !==�Rg  !==�SQR��Tf  +�UT��Vf  -�WV��Xf  *�Yg  *�ZXY��[f  %�\g  %�][\��^f  ++�_g  ++�`^_��af  --�bg  --�cab��df  <<�eg  <<�fde��gf  >>�hg  >>�igh��jf  >>>�kg  >>>�ljk��mf  &�ng  &�omn��pf  |�qg  bor�rpq��sf  ^�tg  ^�ust��vf  !�wg  !�xvw��yf  ~�zg  ~�{yz��|f  &&�}g  &&�~|}��f  ||��g  or������f  ?��g  ?�������f  :��g  colon�������f  =��g  =�������f  +=��g  +=�������f  -=��g  -=�������f  *=��g  *=�������f  %=��g  %=�������f  <<=��g  <<=�������f  >>=��g  >>=�������f  >>>=��g  >>>=�������f  &=��g  &=�������f  |=��g  bor=�������f  ^=��g  ^=�������#&),/258;>ADGJMPSUWZ]`cfilorux{~�������������� .��g  *punctuation*��f  /���I���f  /=���M����� ��g  *div-punctuation*��g  assv-ref��g  
string-ref��g  string-length��g  	substring��f  !bad syntax: character not allowed��g  read-punctuation��g  *eoi*��g  NumericLiteral��f  unexpected right parenthesis��g  srfi-9��� ��g  throw-bad-struct���� ���� ��g  lexical-token-source��g  lexical-token-category��f  unexpected right bracket��f  unexpected right brace��g  reverse!�C 5h�:  �  ]4	
#5 4& >  "  G   '()*        h    �   - 1 3  456�       g  what
			  g  loc			  g  form				  g  args				   g  filenamef   language/ecmascript/tokenize.scm�
	
��				��				��	 		�� 			 	
	  g  nameg  syntax-error� C(R+,-01        h   [   ] L 6S       g  filenamef   language/ecmascript/tokenize.scm�
	#	��		#	,�� 		
   C   h   h   -  1  3 C     `       g  args
			  g  filenamef   language/ecmascript/tokenize.scm�
	#	�� 			


   C4   h@   �   ]!4 54 54 54 O 5� C   �       g  port
		= g  input	&	= g  line		&	= g  column		&	= g  offset		&	=  g  filenamef   language/ecmascript/tokenize.scm�
	
��		 	��	
	!	��		"	��		#	��	&	 	�� 		=  g  nameg  port-source-location� C5R67(89:   hH   4  ]4455$  	64 54455$  
6C   ,      g  delims
		E g  port		E g  loc			E g  token		&	E  g  filenamef   language/ecmascript/tokenize.scm�
	'
��		(	��		(	��		(	��		(	��		)	��		)	��		*	��	$	*	/��	&	*	��	&	*	��	)	+	��	,	+	��	4	+	��	8	+	��	<	,	��	B	,	�� 		E	  g  nameg  
read-until� C;R6<=>?   h@   �   ]	4 5$  C4 5$  C4 5$  C 6   �       g  c
		= g  t		= g  t	*	=  g  filenamef   language/ecmascript/tokenize.scm�
	/
��		0	��		0	��		1	��		1	��	"	2	��	*	1	��	=	3	�� 			=  g  nameg  	char-hex?� C@RA   h   �   ]4 5	0�C   �       g  c
		  g  filenamef   language/ecmascript/tokenize.scm�
	5
��		6	��		6	��		6	�� 		  g  nameg  digit->number� CBR<BAC     h(   �   ]4 5$   6	
44 55	a��C  �       g  c
		&  g  filenamef   language/ecmascript/tokenize.scm�
	8
��		9	��		9	��		:	��		;	��		;	��	!	;	��	#	;	1��	$	;	��	%	;	�� 		&  g  nameg  hex->number� CDRE76HIJK(LMN       h0  %  ]4 >  "  G  4 545$  � C4/5$  4 >  "  G   64*5$  �4 >  "  G  "  c45$  		
64*5$  54 5/�$  4 >  "  G   64 5"���4 5"���4 5"���$  1=�$  4 >  "  G  � C� C 6          g  port
	, g  loc	, g  div?		, g  c1		, g  c		} �  g  filenamef   language/ecmascript/tokenize.scm�
	=
��		?	��		@	��		>	��		B	��	(	A	��	,	D	��	2	D	��	4	E	��	@	A	��	A	F	��	Z	G	��	[	H	��	g	A	��	h	I	��	}	J	��	~	L	
�� �	K	�� �	M	�� �	M	
�� �	N	
�� �	K	�� �	O	�� �	O	�� �	O	
�� �	Q	�� �	R	�� �	S	�� �	S	�� �	U	�� �	U	
�� �	J	�� �	J	�� �	J	�� �	A	�� �	W	�� �	X	��	X	4��	X	 ��	Y	"��#	Y	��,	[	�� +	,	  g  nameg  
read-slash� CORP7JE6Q<RSHT;(UV      hx  %  ]B4/\
5" G4 54/5$  �4 >  "  G  "  �45$  "  T45		$  	"  ;45		$  	"  #4$5		$  	"  	4_5			�$  454	5�
� C4 >  "  G  4 5�"��L4 5"��<4\5$  @4 >  "  G  4 54 544\55"���64 5"���             g  port
	q g  loc	q g  terms		q g  str		\ g  head		\ g  
terminator		\ g  c		A � g  flags		A � g  t		H � g  t			\ � g  t			q � g  t		 � � g  value	 � � g  echar	+R  g  filenamef   language/ecmascript/tokenize.scm�
	]
��		_	��		_	��		`	��		a	��		a	��		c	
��	+	b	��	,	d	
��	A	f	
��	B	g	��	H	g	��	V	h	��	\	h	��	k	i	��	q	h	�� �	j	�� �	h	�� �	k	�� �	h	�� �	g	�� �	m	*�� �	n	*�� �	m	$�� �	l	�� �	l	$�� �	l	�� �	o	�� �	p	�� �	p	,�� �	p	�� �	f	
�� �	f	�� �	f	/��	f	
��	q	
��	b	��	r	
��%	s	��+	s	
��.	t	��9	u	��@	u	(��J	u	��R	t	��V	w	��\	w	
��\	`	��]	`	��i	`	5��q	`	�� 4	q	  g  nameg  read-regexp� CNREP7JHW6<(XY@ZD[R;\        hX  5  ]J4 54\
5" (4 545$  4 >  "  G  � C4\5$ �4 >  "  G  4 5'�$  "  "�$  "  \�$  " sb�$  " df�$  " Un�$  
" Fr�$  " 7t�$  	" (v�$  " 0�$  <4 545$   "  45$  4	
5"   "  �x�$  s4 54 5"  4	455"  =45$  .45$  4	45�45�5"  "���"  "���"  Xu�$  M4 54 54 5	4 5
444	
5	55
	"  4454 55"���	64 5"��� -      g  port
	W g  loc	W g  c			W g  terms		W g  str		F g  
terminator		%F g  c		x g  next	@ g  a	T� g  b	]� g  a	� g  b	� g  c		� g  d	
�	 g  echar	<  g  filenamef   language/ecmascript/tokenize.scm�
	z
��		{	��			{	��		|	��		|	��	 �	��	 �	��	% �	��	( �	��	4 �	
��	5 �	��	J �	 ��	Q �	��	S �	��	_ �	
��	` �	��	r	~	��	x	~	�� �		
��	 �	�� �	�� �	�� �	��# �	��- �	��. �	��2 �	��7 �	��M		
��N �	��T �	��W �	��] �	��d �	��h �	��k �	>��u �	��y �	��z �	��� �	��� �	$��� �	��� �	��� �	)��� �	#��� �	:��� �	 ��� �	���		
��� �	��� �	��� �	��� �	��� �	��� �	��� �	��� �	��� �	��� �	��� �	.�� �	�� �	�� �	�� �	��# �	%��* �	!��6 �	��< �	��@ �	��F �	��F �	��G �	��W �	�� J	W	  g  nameg  read-string� C]R��RR6Q<JS�H(E7  h  �  ]*"  �45$  "  T45$  "  ;45$  "  #4$5$  "  	4_5�$  V445545$  	� C4
5$  
645	� C4 >  "  G  4 5�"��4 5"��       x      g  port
	 g  loc	 g  c		 � g  chars		 � g  t			q g  t		!	m g  t		6	j g  t		M	g g  word	 � � g  t	 � � g  value	 � �  g  filenamef   language/ecmascript/tokenize.scm�
 �
��	 �	��	 �	��	 �	��	 �	��	! �	��	0 �	��	6 �	��	E �	��	M �	��	\ �	��	n �	��	u �	��	v �	��	y �	"�� � �	�� � �	�� � �	�� � �	
�� � �	 �� � �	�� � �	
�� � �	�� � �	�� � �	*�� � �	�� � �	*�� � �	�� � �	�� � �	�� � �	$�� � �	�� � �	�� � �	�� � �	'�� �	�� %		  g  nameg  read-identifier� CRJ7E6B<(@D   h�  /  ]:44 5.5$  0"  4 54 545$  6" ?" *45$  C45$  .4 >  "  G  4 5	
�45�"���4e5$  "  	4E5$ 4 >  "  G  4 545$  45"  f4+5$  4 >  "  G  	"  A4-5$  4 >  "  G  
"  45$  	"  45"  k45$  "  45$  24 >  "  G  4 54	
�455"���
�$  	�"  4	
5�C4 5
"���4.5$  �4 >  "  G  "  d45$  "  45$  :4 >  "  G  4 5454	
5���"����"���4 5	�"���C45"���405$ Y4x5$  "  	4X5$  �4 >  "  G  4 545$  "  4>  "  G  "  <45$  .4 >  "  G  4 5	�45�"���C
"���45$  �"  �45$  C45$  n485$  "  	495$  4>  "  G  "   4 >  "  G  4 5	�45�"��yC
"��k"��["��W      '      g  port
	� g  loc	� g  c0		!� g  c1		*� g  c1		Eo g  acc		Eo g  t	 � � g  c	 �O g  add	O� g  c	U� g  e	U� g  c	�W g  dec	�W g  n	�W g  t	�� g  c	�> g  c	�3 g  acc	�3 g  c	M� g  acc	M� g  t	o�  g  filenamef   language/ecmascript/tokenize.scm�

��		��		��		��		��		��	!	��	$	��	*	��	-	��	7	��	=	��	E"	��	F$	
��	P#	��	T%	
��	^#	��	_&	
��	q'	��	|(	��	}(	�� �(	�� �'	
�� �)	�� �)	
�� �)	�� �#	�� �*	
�� �+	�� �+	�� �,	�� �,	�� �-	�� �-	,�� �-	�� �/	�� �,	�� �/	-��0	��,	��0	-��11	��;,	��B3	��F3	,��L3	��O+	
��U5	��V6	��`6	��f6	0��p6	��q7	���8	���8	*���8	/���8	8���8	*���8	���:	���:	���:	1���:	*���:	;���:	���5	���5	���5	���;	
���#	���<	
���=	
���>	���>	��>	.��>	��?	��!@	��*A	"��1A	4��:A	��;A	��>B	��H@	��OE	��WE	��W=	
��X=	��`=	.��l=	
��o"	��r"	���"	���	���	���	���	���	���	/���	���		���
	���
	���	���	���	���	���	���	���	��	
��	��	�� 	��!	#��(	��0	��3	��?	��I	��M	��N	��X	��\	��f	��g	��o	��}	&���	���	���	!���	���	���	���	���	 ���	���	���	�� �	�	  g  nameg  read-numeric� C R��R��R����       h�   �  ]
(   C4 4��
55$  D4��5�$  ���"  4L �4��5��� 5� � "���4��
5  � "��     �      g  nodes
	 � g  puncs	 � g  t		 �  g  filenamef   language/ecmascript/tokenize.scm�
	��	�	��	�	��	�	,��	�	8��	�	,��	�	��	�	��	%�	(��	*�	7��	-�	(��	/�	%��	3�	!��	8�	9��	:�	%��	A�	/��	F�	3��	G�	7��	L�	B��	P�	7��	S�	9��	U�	3��	Z�	/��	[�	%��	`�	+��	h�	!��	i�	,��	n�	8��	r�	,��	v�	&��	y�	 �� ��	�� !	 �	  g  nameg  lp� CO   Q  4 �i5  �E7H(�      hp   F  ]""  U45$  )4 >  "  G  4 5��"���$  � C64 5L "��� >      g  port
		o g  loc		o g  c			[ g  tree			[ g  	candidate			[ g  t			[  g  filenamef   language/ecmascript/tokenize.scm�
�	��	�	��	�	
��	�	��	�	��	*�	��	3�	$��	6�	4��	@�	��	F�	��	P�	
��	U�	��	[�	
��	[�	��	\�	��	o�	�� 		o	   C O  �R75EO]6�QJ< H��     hP    ]"4 54 5	�$  "  /�$  "  !�$  "   �$  "  ��$  4 >  "  G   6
�$  "  �$  4 >  "  G   6/�$  
 6"�$  "  '�$   645$  C4	5$  "  "4
$5$  "  	4
_5$   645$  4 5� C 6       g  port
	O g  div?	O g  c		O g  loc		O g  t	 � g  t	 � g  value	9G  g  filenamef   language/ecmascript/tokenize.scm�
�
��	�	��	
�	��	�	��	�	��	V�	��	o�	��	x�	�� ��	�� ��	�� ��	�� ��	�� ��	�� ��	�� ��		�� ��	�� ��		�� ��	�� ��		�� ��	�� ��		���	���	��%�		��&�		��0�	��1�	1��9�		��?�	��F�		��O�		��  	O	  g  nameg  
next-token� CRH�W    hX     ]4L M5  �$  : �&  , 
�&  "  &  "  �"  "  N C    �       g  tok
		T g  cat	!	E  g  filenamef   language/ecmascript/tokenize.scm�
�	��	�	��	�	��	�	��	�	��	�	��	!�	#��	!�	��	'�	'��	+�	��	,�	��	4�	'��	8�	��	9�	��	A�	'��	B�	��	Q�	�� 		T
   C       h   �   ]	H O C �       g  port
		 g  div?		  g  filenamef   language/ecmascript/tokenize.scm�
�
��	�	�� 		  g  nameg  make-tokenizer� C R�H(+(����.1�"%�7�W     hh  �  ]M$  C4LM5  �$   �&   
�"   "   �$   M �N " ��$  }"  14 �&   �"  	4 	5>  "  KG  "  DM �$  9M ��&  
�"  	4
5&  	M �N "  "���"  "���" F�$   M �N " 2�$  }"  14 �&   �"  	4 	5>  "  KG  "  DM �$  9M ��&  
�"  	4
5&  	M �N "  "���"  "���"  ��$   M �N "  ��$  }"  14 �&   �"  	4 	5>  "  KG  "  DM �$  9M ��&  
�"  	4
5&  	M �N "  "���"  "���"  �$  	M �N"    �$  : �&  , 
�&  "  &  "  �"  "  N C      �      g  tok
	b g  key	6 g  s	 � � g  s	.M g  s	�� g  cat/S  g  filenamef   language/ecmascript/tokenize.scm�
�	��	�	��	
�	
��	�	��	�	
��	�	��	�	��	 �	��	*�	+��	6�	��	F�	��	H�	��	U�	��	Z�	��	^�	!��	a�	!��	~�	�� ��	�� ��	�� ��	�� ��	5�� ��	�� ��	B�� ��	�� ��	�� ��	�� ��	�� ��	�� ��	�� ��	�� ��	�� ��	!�� ��	!���	��$�	��'�	��+�	��.�	5��.�	��O�	B��S�	��V�	��X�	��u�	��z�	��|�	����	����	����	!����	!����	����	����	����	����	5����	����	B����	����	����	���	���	���	���	��"�	��%�	��/�	'��/�	��5�	+��9�	��:�	"��B�	+��F�	��G�	"��O�	+��P�	"��_�	�� N	b
   C       h    �   ]HHH O C       �       g  port
		 g  div?		 g  eoi?			 g  stack			  g  filenamef   language/ecmascript/tokenize.scm�
�
��	�	��	�	�� 		  g  nameg  make-tokenizer/1� C!R ��      h8   #  ]4 5"   45 &  6�"���"���        g  port
		6 g  next			6 g  out			/ g  tok			/  g  filenamef   language/ecmascript/tokenize.scm�
�
��	�	��		�	��	�	��	�	��	�	��	�	��	�	��	$�	��	)�	��	/�	��	/�	��	0�	��	6�	�� 		6  g  nameg  tokenize� CR!��      h8   %  ]4 5"   45 &  6�"���"���        g  port
		6 g  next			6 g  out			/ g  tok			/  g  filenamef   language/ecmascript/tokenize.scm�
�
��	�	��		�	��	�	��	�	��	�	��	�	��	�	��	$�	��	)�	��	/�	��	/�	��	0�	��	6�	�� 		6  g  nameg  
tokenize/1� C"RC      �      g  m
		, g  lp
(�(� g  	punc-tree
(�*�  g  filenamef   language/ecmascript/tokenize.scm�		
��9	
���	
��'	'
��O	/
��	5
��	8
���	=
��`	]
��(	z
��* �	��- �
��0 �	��4 �
��� �
��&S
��&VJ	��&ZI
��&]z	��&ay
��(�	��(�	#��(�	��(�	��*�~
��.S�
��0��
��7��
��9^�
��:��
�� 	:�
   C6 