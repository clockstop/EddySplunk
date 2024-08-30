import sys
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

def plot_invocation_times(filename):
    # Load data
    df = pd.read_csv(filename, sep=':', header=None, names=['Invocation', 'Time'])
    
    # Clean up the 'Time' column
    df['Time'] = pd.to_numeric(df['Time'].str.strip().str.replace(' ms', ''), errors='coerce')
    
    # Calculate statistics
    mean = df['Time'].mean()
    p90 = np.percentile(df['Time'].dropna(), 90)
    p95 = np.percentile(df['Time'].dropna(), 95)
    p99 = np.percentile(df['Time'].dropna(), 99)
    
    # Plot the data
    plt.figure(figsize=(10, 5))
    plt.plot(df['Time'], marker='o', linestyle='-', color='b', label='Invocation Times')
    
    # Add lines for mean, p95, and p99
    plt.axhline(mean, color='r', linestyle='--', linewidth=1, label=f'Mean: {mean:.2f} ms')
    plt.axhline(p90, color='b', linestyle='--', linewidth=1, label=f'P90: {p90:.2f} ms')
    plt.axhline(p95, color='g', linestyle='--', linewidth=1, label=f'P95: {p95:.2f} ms')
    plt.axhline(p99, color='y', linestyle='--', linewidth=1, label=f'P99: {p99:.2f} ms')
    
    plt.xlabel('Invocation Number')
    plt.ylabel('Time (ms)')
    plt.title('Lambda Invocation Times')
    plt.legend()
    plt.grid(True)
    plt.savefig('invocation_times.png')
    plt.show()

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python visualize_times.py <filename>")
        sys.exit(1)
    
    filename = sys.argv[1]
    plot_invocation_times(filename)
