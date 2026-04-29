# What tables do we have and why?

1. Team
2. Engineer
3. Schedule
4. EngineerAndSchedule
5. Override

Goal: Which engineer is on call at a given time

Working backwards: "For a given time, what schedule is running and what team is assigned to that schedule and which engineer on that team is on-call at that time on that schedule"

Analysing this:
- Look at the nouns involved: Schedule, Team, Engineer
- Look at the relationships:
    - 1 Engineer maps to 1 Team
    - 1 Team maps to many Engineers
    - 1 Schedule maps to 1 Team
    - 1 Team maps to many Schedules
    - 1 Engineer maps to many Schedules
    - 1 Schedule maps to many Engineers
- Check if the relationships carry any of their own attributes
    - An Engineer is on a Schedule from a beginning time to an end time
    - "from a beginning time to and end time" needs to live somewhere 
- Check edge cases
    - Override the schedule to put this new Engineer on-call from begining time to end time

This give us a clear picture of the full world of possibilities and relationships we care about

From this we create:
- The 3 nouns as tables: Schedule, Team, Engineer
- The Engineer / Schedule relation as a table: EngineerAndSchedule
- The edge case as a table: Override

TABLE TEAM
* ID - INTEGER - Primary key
* Name - STRING

TABLE ENGINEER
* ID - INTEGER - Primary key
* Name - STRING
* Team - INTEGER - Foreign Key

TABLE SCHEDULE
* ID - INTEGER - primary key
* Begin - TIMESTAMP
* End - TIMESTAMP
* ShiftLength - INTEGER 
* TeamID - INTEGER - Foreign Key

TABLE EngineerAndSchedule
* Engineer - INTEGER - foreign key
* Schedule - INTEGER - foreign key
* RotationPosition - INTEGER

TABLE OVERRIDE
* ID - INTEGER - primary key
* Schedule - INTEGER - foreign key
* Begin - TIMESTAMP
* End - TIMESTAMP
* NewEng - INTEGER