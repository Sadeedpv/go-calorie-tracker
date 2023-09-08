import { create } from "zustand"
import { calorieData, totalCalories } from "./type"

interface calorieState{
    calorieData: calorieData,
    totalCalories: totalCalories,
    fetchCalorieData: () => void,
    fetchTotalCalories: () => void
}

export const useCalorieStore = create<calorieState>()((set) => ({
    calorieData: [],
    totalCalories: 0,
    fetchCalorieData: () => {
        fetch(`${import.meta.env.VITE_PORT}/calories`,
            {
                method: "GET",
                headers: {
                    "Content-Type": "application/json"
                }
            }
        )
        .then(res => res.json())
        .then(data => {
            set({ calorieData: data })
        })
    },
    fetchTotalCalories: () => {
        fetch(`${import.meta.env.VITE_PORT}/totalcalories`,
            {
                method: "GET",
                headers: {
                    'Content-Type': 'application/json'
                }
            }
        )
        .then(res => res.json())
        .then(data => {
            set({ totalCalories: data.TotalCalories });
        })
        
    }
}))