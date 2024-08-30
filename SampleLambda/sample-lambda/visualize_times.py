import sys
import pandas as pd
import matplotlib.pyplot as plt

def plot_invocation_times(filename):
    # Load data
    df = pd.read_csv(filename, sep=':', header=None, names=['Invocation', 'Time'])
    
    # Clean up the 'Time' column
    df['Time'] = pd.to_numeric(df['Time'].str.strip().str.replace(' ms', ''), errors='coerce')
    
    # Plot the data
    plt.figure(figsize=(10, 5))
    plt.plot(df['Time'], marker='o', linestyle='-', color='b')
    plt.xlabel('Invocation Number')
    plt.ylabel('Time (ms)')
    plt.title('Lambda Invocation Times')
    plt.grid(True)
    plt.savefig('invocation_times.png')
    plt.show()

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python visualize_times.py <filename>")
        sys.exit(1)
    
    filename = sys.argv[1]
    plot_invocation_times(filename)

