GRANT CONNECT ON DATABASE mtr TO mtr_w;
GRANT USAGE ON SCHEMA field TO mtr_w;
GRANT ALL ON ALL TABLES IN SCHEMA field TO mtr_w;
GRANT ALL ON ALL SEQUENCES IN SCHEMA field TO mtr_w;

--  TODO read only user?