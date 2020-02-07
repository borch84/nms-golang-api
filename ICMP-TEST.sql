--delete from kis.icmp where server='stnp1002b.ussgv008.neustar.com'

insert into kis.icmp (ELEMENT_ID, SERVER, TIMEOUT, NUMBEROFPINGS, PACKETINTERVAL, PACKETSIZE, TYPEOFSERVICE, RETRIES, POLL, FAILURERETESTS, RETESTINTERVAL, DESCRIPTION, HOSTNAMELOOKUPPREFERENCE) 
values ('b22a322f-090e-494a-93a5-4a37973d5507','stnp1002b.ussgv008.neustar.com', 10, 5, 1, 64, 0, 0, 300, 0, 10, 'ICMP stnp1002b.ussgv008.neustar.com element.', 'Default');
insert into kis.icmp (ELEMENT_ID, SERVER, TIMEOUT, NUMBEROFPINGS, PACKETINTERVAL, PACKETSIZE, TYPEOFSERVICE, RETRIES, POLL, FAILURERETESTS, RETESTINTERVAL, DESCRIPTION, HOSTNAMELOOKUPPREFERENCE) 
values ('02f9a7e2-e7dd-43f1-8320-dbadde4575b5', 'stnp1002b.ussgv008.neustar.com', 10, 5, 1, 64, 0, 0, 300, 0, 10, 'ICMP stnp1002b.ussgv008.neustar.com element', 'Default');

