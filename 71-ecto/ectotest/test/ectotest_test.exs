defmodule EctotestTest do
  use ExUnit.Case
  doctest Ectotest

  test "greets the world" do
    assert Ectotest.hello() == :world
  end
end
