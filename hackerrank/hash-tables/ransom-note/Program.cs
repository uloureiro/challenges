using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

class Solution {
    static void Main(String[] args) {
        Console.WriteLine("Enter the number of words available in the magazine and the number of words for the note (use space to split)");
        string[] tokens_m = Console.ReadLine().Split(' ');
        int m = Convert.ToInt32(tokens_m[0]);
        int n = Convert.ToInt32(tokens_m[1]);
        Console.WriteLine("Enter the magazine's words");
        string[] magazine = Console.ReadLine().Split(' ');
        Console.WriteLine("Enter the note's words");
        string[] ransom = Console.ReadLine().Split(' ');
        
        Console.Write(IsMagazineValidToRansom(m,n,magazine,ransom));
    }
    
    private static string IsMagazineValidToRansom(int m, int n, string[] magazine, string[] ransom)
    {
        // the magazine hasn't enough words
        if(m < n) return "No";
        
        Dictionary<string, int> magazineDict = magazine.GroupBy(a => a).ToDictionary(a => a.Key, a => a.Count());
        
        foreach(var word in ransom) 
        {
            if(magazineDict.ContainsKey(word)) // Do I have this word in the mag.?
            {
                if(magazineDict[word] > 0) // Do I have enough words available?
                {
                    magazineDict[word]--;
                }
                else
                {
                    return "No"; // I do not have any of this word available... no ransom
                }
            }
            else 
            {
                return "No"; // I do not have this word in the mag.
            }
        }
        
        return "Yes";
    }
}
