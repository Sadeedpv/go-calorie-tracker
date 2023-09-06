import Table from 'react-bootstrap/Table'
import React from 'react'


type calorieData = {
  id: number,
  food: string,
  calorie: number
}[]

type totalCalories = number

function StripedTable() {
  const [calorieData, setCalorieData] = React.useState<calorieData|null>();
  const [totalCalories, settotalCalories] = React.useState<totalCalories|null>();
  React.useEffect(() => {
    fetch(`${import.meta.env.VITE_PORT}/calories`)
      .then((res) => res.json())
      .then((data) => setCalorieData(data));
    fetch(`${import.meta.env.VITE_PORT}/totalcalories`)
      .then((res) => res.json())
      .then((data) => settotalCalories(data.TotalCalories))
    
  }, [])
  return (
    <Table striped bordered hover>
      <thead>
        <tr>
          <th>ID</th>
          <th>FOOD</th>
          <th>CALORIES</th>
        </tr>
      </thead>
      <tbody>
        {
          calorieData?.map((data, index) => {
            return (
              <tr key={index}>
                <td>{data.id}</td>
                <td>{data.food}</td>
                <td>{data.calorie}</td>
              </tr>
            )
          })
        }
        <tr>
          <td colSpan={2} className='endRow'>Total Calories</td>
          <td>{totalCalories}</td>
        </tr>
      </tbody>
    </Table>
  );
}

export default StripedTable