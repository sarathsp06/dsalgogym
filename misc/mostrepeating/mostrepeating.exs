#!/usr/bin/env elixir
defmodule Script do
  def main(args) do
    IO.puts(mostrepeating(args))
  end

  def mostrepeating([]), do: :ok
  def mostrepeating(args) do
	  args 
    |> Enum.map(fn(s)->String.to_integer(s)end) 
    |> Enum.sort() 
    |> (fn(s)->Enum.at(s,div(length(s),2)) end).()
  end
end

Script.main(System.argv)