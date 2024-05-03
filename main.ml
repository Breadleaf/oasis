let () =
  let pattern = "hello" in
  let text = "hello world" in
  if Rm.regex_match pattern text then
    print_endline "Matched"
  else
    print_endline "Not matched"

let () =
  let random_number = Ru.generate_random_number () in
  print_endline (string_of_int random_number)
