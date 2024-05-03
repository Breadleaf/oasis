let regex_match pattern text =
  let regex = Re.compile (Re.Perl.re pattern) in
  Re.execp regex text
