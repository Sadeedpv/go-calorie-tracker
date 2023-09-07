import Table from 'react-bootstrap/Table'
import { calorieData, totalCalories } from '../utils/type';



function StripedTable({calorieData, totalCalories}:{calorieData:calorieData,totalCalories:totalCalories}) :JSX.Element{
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