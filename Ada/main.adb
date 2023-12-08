with Ada.Text_IO;

procedure CheckPrime is
   function Is_Prime (Num : Integer) return Boolean is
   begin
      if Num < 2 then
         return False;
      end if;

      for I in 2 .. Num - 1 loop
         if Num mod I = 0 then
            return False;
         end if;
      end loop;

      return True;
   end Is_Prime;

   procedure Generate_CheckPrime_File is
      Output : Ada.Text_IO.File_Type;
   begin
      Ada.Text_IO.Create (Output, Ada.Text_IO.Out_File, "checkPrime.adb");

      Ada.Text_IO.Put_Line (Output, "with Ada.Text_IO;");
      Ada.Text_IO.Put_Line (Output, "procedure CheckPrime is");
      Ada.Text_IO.Put_Line (Output, "   function Check_Prime (Num : Integer) return Boolean is");
      Ada.Text_IO.Put_Line (Output, "   begin");

      for I in 1 .. 1000 loop
         Ada.Text_IO.Put_Line (Output, "      if Num = " & Integer'Image (I) & " then");
         Ada.Text_IO.Put_Line (Output, "         return " & Boolean'Image (Is_Prime (I)) & ";");
         Ada.Text_IO.Put_Line (Output, "      end if;");
      end loop;

      Ada.Text_IO.Put_Line (Output, "      return False;");
      Ada.Text_IO.Put_Line (Output, "   end Check_Prime;");
      Ada.Text_IO.Put_Line (Output, "begin");
      Ada.Text_IO.Put_Line (Output, "   null;");
      Ada.Text_IO.Put_Line (Output, "end CheckPrime;");

      Ada.Text_IO.Close (Output);
   end Generate_CheckPrime_File;

begin
   Generate_CheckPrime_File;
end CheckPrime;
