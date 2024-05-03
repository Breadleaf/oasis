let regex_match pattern text =
  let regex = Re.compile (Re.Perl.re pattern) in
  Re.execp regex text

let random_match text =
  regex_match (string_of_int (Ru.generate_random_number ())) text
