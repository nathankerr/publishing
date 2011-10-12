\NeedsTeXFormat{LaTeX2e}[1995/12/01]
\ProvidesClass{article}[2011/10/12]
\LoadClass[article,a5paper,twoside,openright]{memoir}

% Growth based on http://en.wikipedia.org/wiki/Preferred_number R10 series
% one to two alphabets = 18.7pt - 9.35pt
% basesizes in series: 10pt, 12.59pt, 15.85pt
% Numbers increased by a mulitple of the 10th root of 10 (1.26) and then rounded
% Lineheight is 1.26*fontsize
% skips are 1/3 of vspace remaining for heading
% (2*\baselineheight - lineheight)/3
\newcommand\@sizeminiscule{3.479pt} %5
\newcommand\@sizetiny{4.379pt} %6
\newcommand\@sizescriptsize{5.513pt} %7
\newcommand\@sizefootnotesize{6.941pt} %8
\newcommand\@sizesmall{8.738pt} %9
\newcommand\@sizenormal{11pt} % 1.1 * 10^(10/10)
\newcommand\@sizelarge{13.848pt} % 11 * 10^(1/10)
\newcommand\@sizeLarge{17.434pt} % 2
\newcommand\@sizeLARGE{21.948pt} % 3
\newcommand\@sizehuge{27.631pt} % 4
\newcommand\@sizeHuge{34.785pt} % 5
\newcommand\@sizeHUGE{43.792pt} % 6
\newcommand\@sizeHUGER{55.131pt} % 7


% Packages
\RequirePackage{url}

% Fonts and Titles
\usepackage{fontspec}
\setmainfont[Ligatures=TeX]{Georgia}
\setsansfont[Ligatures=TeX]{Verdana}
\setmonofont[Ligatures=TeX]{Andale Mono}
\setlength\parindent{0pt}
\setlength\parskip{\@sizefootnotesize}

% Font Sizes
\renewcommand\miniscule{
	\fontsize{\@sizeminiscule}{\@sizetiny}
	\selectfont
}
\renewcommand\tiny{
	\fontsize{\@sizetiny}{\@sizescriptsize}
	\selectfont
}
\renewcommand\scriptsize{
	\fontsize{\@sizescriptsize}{\@sizefootnotesize}
	\selectfont
}
\renewcommand\footnotesize{
	\fontsize{\@sizefootnotesize}{\@sizesmall}
	\selectfont
}
\renewcommand\small{
	\fontsize{\@sizesmall}{\@sizenormal}
	\selectfont
}
\renewcommand\normalsize{
	\fontsize{\@sizenormal}{\@sizelarge}
	\selectfont
}
\renewcommand\large{
	\fontsize{\@sizelarge}{\@sizeLarge}
	\selectfont
}
\renewcommand\Large{
	\fontsize{\@sizeLarge}{\@sizeLARGE}
	\selectfont
}
\renewcommand\LARGE{
	\fontsize{\@sizeLARGE}{\@sizehuge}
	\selectfont
}
\renewcommand\huge{
	\fontsize{\@sizehuge}{\@sizeHuge}
	\selectfont
}
\renewcommand\Huge{
	\fontsize{\@sizeHuge}{\@sizeHUGE}
	\selectfont
}
\renewcommand\HUGE{
	\fontsize{\@sizeHUGE}{\@sizeHUGER}
	\selectfont
}

% Page Layout
%\setbinding{9mm}
\setlxvchars[\normalsize]
\settypeblocksize{*}{\lxvchars}{1.414}
\setulmargins{*}{*}{1.414}
\setlrmargins{*}{*}{1.414}
\checkandfixthelayout[lines]

% Headers and footers
\pagestyle{plain} %which also solves the style used by chapters
\makeevenfoot{plain}{\thepage}{}{} % style left center right
\makeoddfoot{plain}{}{}{\thepage} % style left center right

% Title
\renewcommand\@maketitle{
	\thispagestyle{empty}
	\vspace{-\parskip}
	\@author \hfill \@date
	\vspace{2pt}
	\hrule
	\begin{center}
		\vspace{-\@sizelarge}
		\LARGE
		\@title
	\end{center}
}

\setcounter{secnumdepth}{5}
% Heading 1 (chapter)
\renewcommand{\chapnamefont}{\LARGE}
\renewcommand{\chapnumfont}{\LARGE}
\renewcommand{\chaptitlefont}{\LARGE}
\renewcommand\printchapternonum{\vphantom{\chapnamefont Chapter \chapnumfont 2}\vspace{2pt}}
\renewcommand{\printchaptertitle}[1]{#1}
\renewcommand\printchaptername{}
\renewcommand\printchapternonum{}
\setlength\afterchapskip{-\parskip}

% Heading 2 (section)
\setbeforesecskip{\@sizelarge}
\setsecindent{0pt}
\setsecheadstyle{\Large}
\setaftersecskip{1pt}

% Heading 3 (subsection)
\setbeforesubsecskip{\@sizenormal}
\setsubsecindent{0pt}
\setsubsecheadstyle{\large}
\setaftersubsecskip{1pt}

% Heading 4 (subsubsection)
\setbeforesubsubsecskip{\@sizenormal}
\setsubsecindent{0pt}
\setsubsubsecheadstyle{\normalsize}
\setaftersubsubsecskip{1pt}

% Heading 5 (paragraph)
\setbeforeparaskip{\parskip}
\setparaindent{1pt}
\setparaheadstyle{\normalsize}
\setafterparaskip{-1em}

% Heading 6 (subparagraph)
\setbeforesubparaskip{\parskip}
\setsubparaindent{1pt}
\setsubparaheadstyle{\normalsize}
\setaftersubparaskip{-1em}
