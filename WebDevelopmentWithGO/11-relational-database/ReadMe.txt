The large tech firms started developing technology to deal with storing data at sacle because relational databases they can scale vertically, you can increate the size of the machine, increate the proccessing power, increase the RAM. They can scale vertically ok, but the don't scale horizontally wel , it is not easy to take relational database and store it across many machines and that has to do with a couple reasons but mostly beacuse you want to have consistency of transcations and to do that across many machine with relational database is hard.
We are going to use MySql as relational database.

to be able to use the MySql with Go code, we need to get the driver for MySql:
	> go get github.com/go-sql-driver/mysql