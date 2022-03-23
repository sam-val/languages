using System;
using System.Collections;

class Pyramid
{
	static bool reverse = false;
	static int length = 5;
	public static void Main(string[] args)
	{
		if (args.Length == 0)
		{
			DisplayPyramid(length, reverse);		
		}
		else
		{
			if (args.Length > 0)
			{
				try 
				{
					int number;
					if (args.Length == 1)
					{
						if(args[0] == "--reverse" | args[0] == "-r")
						{
							reverse = true;
						} else if (Int32.TryParse(args[0], out number))
						{
							length = number;
						}
						else 
						{
							throw new Exception("invalid argument");
						}
						
					}
					if (args.Length == 2)
					{
						int input_length = Int32.Parse(args[0]);
						length = input_length;
						

						if(args[1] == "--reverse" | args[1] == "-r")
						{
							reverse = true;
						}
						else 
						{
							throw new Exception("invalid argument");
						}
					
					}	
				} catch (Exception e)
				{
					Console.WriteLine("invalid argument");
					Console.WriteLine(e.GetType().Name);
					return;
				}
						
					
			}
			else
			{
				Console.WriteLine("error");	
			}
			DisplayPyramid(length, reverse);
		}
	}
	
	public static void DisplayPyramid(int length, bool mode)
	{
		int lengthLastLine = length * 2 - 1;
		string val = "#";
		string space = " ";
		if (mode)
		{
			for (int i=length-1; i >= 0; i--)
			{
				for (int j = lengthLastLine-1; j >= 0; j--)
				{
					if ((j <= (lengthLastLine/2 + i)) && (j >= (lengthLastLine/2 - i)))
					{
						Console.Write(val);
					}
					else 
					{
						Console.Write(space);	
					}
				}
				Console.WriteLine("");

			}
		}
		else 
		{
			for (int i=0; i < length; i++)
			{
				for (int j = 0; j < lengthLastLine; j++)
				{
					if ((j <= (lengthLastLine/2 + i)) && (j >= (lengthLastLine/2 - i)))
					{
						Console.Write(val);
					}
					else 
					{
						Console.Write(space);	
					}
				}
				Console.WriteLine("");

			}
			Console.Write($"normal pyramid, length {length}");	
		}
	}
}
