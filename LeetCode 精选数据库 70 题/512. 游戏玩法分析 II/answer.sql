select player_id,device_id
from Activity
where (player_id,event_date) in (select player_id,min(event_date) as event_date
                                 from Activity
                                 group by player_id)



select
    player_id, device_id
from
    (
        select
            player_id,
            device_id,
            dense_rank() over(partition by player_id
                          order by event_date asc) rnk
        from activity
    ) a where rnk=1