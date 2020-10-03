defmodule ExampleTest do
  use ExUnit.Case
  doctest Script

  test "[1,2,3,1,1,1,1,1,3,4,1,1,2]" do
    input = [1,2,3,1,1,1,1,1,3,4,1,1,2] |> 
      Enum.map(fn(x)->Integer.to_string(x)end)
    assert Script.mostrepeating(input) == 1
  end
  test "[]" do
    assert Script.mostrepeating([]) == :ok
  end
end