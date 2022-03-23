using System;
using System.Collections.Generic;

class @RPS 
{
	static int userPoints = 0, comPoints = 0, round = 1;
	

	public static int Compute(int input_n, int computerChoice)
	{
			if (input_n == computerChoice)
			{
				return 0;
			}
			else if (input_n == 1)
			{
				if (computerChoice == 2)
				{
					return -1;

				}	
				else 
				{
					return 1;
				}
			}
			else if (input_n == 2)
			{
				if (computerChoice == 1)
				{
					return 1;
				}	
				else 
				{
					return -1;
				}
			}
				
			else 
			{
				if (computerChoice == 1)
				{
					return -1;
				}	
				else
				{
					return 1;
				}
			}
	}
	static void Main(string[] args)
	{
		Console.WriteLine("Starting Rock, Paper, Scissors game...");
		var options = new Dictionary<int,string>
		{
			{1, "Rock"},
			{2, "Paper"},
			{3, "Scissors"}
		};
		while (round <= 3)
		{
			Console.Write("Choose options: \n1. Rock\n2. Paper\n3. Scissors.\nEnter: ");
			bool valid = false;
			int input_n = 0;
			
			while (!valid)
			{
				var input = Console.ReadLine();
				try 
				{
					input_n = Convert.ToInt32(input);
					if (!options.ContainsKey(input_n))
					{
						throw new Exception("Please choose between 1,2 and 3");
					}
					valid = true;
				}
				catch (Exception e)
				{
					Console.Write("Please choose 1, 2, or 3.\nEnter: ");
				}
				
			}
			Console.WriteLine("\n---Player played " + options[input_n]);
			Random rnd = new Random();
			int computerChoice =  rnd.Next(1,4);
			Console.WriteLine("---Computer played " + options[computerChoice]);
			
			// check winning or losing:
			int rs = Compute(input_n,computerChoice);
			string info = "";
			if (rs == 0)
			{
				info ="round #"+round + " It's a draw!";
			}else if (rs == 1)
			{
				info = "Player won #" + round + " round";
				userPoints++;
				
			}else if (rs == -1)
			{
				info = "Computer won #" + round + " round";
				comPoints++;
			}
			Console.WriteLine(">>>"+info + "\n");	
			round++;
		}
		if (userPoints > comPoints)
		{
			Console.WriteLine("YOU WON!!!");
		}else if (userPoints == comPoints)
		{
			Console.WriteLine("IT'S A DRAW. YEYYYY!!!");
		} 
		else
		{
			Console.WriteLine("YOU LOST");
		}
	}
}
