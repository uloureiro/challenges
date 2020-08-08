using System;
using System.Collections;
using System.Collections.Generic;

class Solution
{

  static void Main(String[] args)
  {
    int t = Convert.ToInt32(Console.ReadLine());
    for (int a0 = 0; a0 < t; a0++)
    {
      string expression = Console.ReadLine();
      Console.WriteLine(ValidateBracketsBalance(expression));
    }
  }

  public static string ValidateBracketsBalance(string expression)
  {
    if (String.IsNullOrEmpty(expression)) // the number of elements is zero
      return "NO";

    char[] expressionArray = expression.ToCharArray();
    if (expressionArray.Length % 2 > 0) // the number of elements is odd, the expression cannot be balanced
      return "NO";

    Dictionary<char, char> brackets = new Dictionary<char, char> { { ')', '(' }, { ']', '[' }, { '}', '{' } };
    Stack openingBracketStack = new Stack();

    for (var i = 0; i < expressionArray.Length; i++)
    {
      // stacks the element if it is an opening bracket
      if (expressionArray[i] == '(' || expressionArray[i] == '[' || expressionArray[i] == '{')
        openingBracketStack.Push(expressionArray[i]);
      else
      {
        // if the current closing element do not match to the current top element from the stack, then it's unbalanced
        if (openingBracketStack.SafePeek() == null || brackets[expressionArray[i]] != (char)openingBracketStack.Peek())
          return "NO";
        else
          openingBracketStack.SafePop();
      }
    }

    if (openingBracketStack.Count > 0) return "NO"; // if anything has left in the stack then it's unbalanced
    else return "YES";
  }
};

public static class ExtensionMethods
{

  public static object SafePeek(this Stack stack)
  {
    if (stack.Count > 0) return stack.Peek();
    return null;
  }

  public static object SafePop(this Stack stack)
  {
    if (stack.Count > 0) return stack.Pop();
    return null;
  }

}