DO $$
declare phonenumber varchar := '12345678';
BEGIN
for cnt in 1..10 loop            
                insert into  phone (id, phone) VALUES (cnt, phonenumber);
   end loop;
END;
$$;